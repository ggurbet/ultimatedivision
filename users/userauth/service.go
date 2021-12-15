// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package userauth

import (
	"context"
	"crypto/subtle"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/zeebo/errs"
	"golang.org/x/crypto/bcrypt"

	"ultimatedivision/console/emails"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
	"ultimatedivision/pkg/cryptoutils"
	"ultimatedivision/users"
)

const (
	// TokenExpirationTime after passing this time token expires.
	TokenExpirationTime = 24 * time.Hour
	// PreAuthTokenExpirationTime after passing this time token expires.
	PreAuthTokenExpirationTime = 2 * time.Hour
)

var (
	// ErrUnauthenticated should be returned when user performs unauthenticated action.
	ErrUnauthenticated = errs.Class("user unauthenticated error")

	// Error is a error class for internal auth errors.
	Error = errs.Class("user auth internal error")

	// ErrPermission should be returned when user permission denied.
	ErrPermission = errs.Class("permission denied")

	// ErrAddressAlreadyInUse should be returned when users email address is already in use.
	ErrAddressAlreadyInUse = errs.Class("email address is already in use")
)

// Service is handling all user authentication logic.
//
// architecture: Service
type Service struct {
	users        users.DB
	signer       auth.TokenSigner
	emailService *emails.Service
	log          logger.Logger
}

// NewService is a constructor for user auth service.
func NewService(users users.DB, signer auth.TokenSigner, emails *emails.Service, log logger.Logger) *Service {
	return &Service{
		users:        users,
		signer:       signer,
		emailService: emails,
		log:          log,
	}
}

// Token authenticates User by credentials and returns auth token.
func (service *Service) Token(ctx context.Context, email string, password string) (token string, err error) {
	user, err := service.users.GetByEmail(ctx, email)
	if err != nil {
		return "", Error.Wrap(err)
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		return "", ErrUnauthenticated.Wrap(err)
	}

	claims := auth.Claims{
		UserID:    user.ID,
		Email:     user.Email,
		ExpiresAt: time.Now().UTC().Add(TokenExpirationTime),
	}

	token, err = service.signer.CreateToken(ctx, &claims)
	if err != nil {
		return "", Error.Wrap(err)
	}

	return token, nil
}

// LoginToken authenticates user by credentials and returns login token.
func (service *Service) LoginToken(ctx context.Context, email string, password string) (token string, err error) {
	user, err := service.users.GetByEmail(ctx, email)
	if err != nil {
		return "", Error.Wrap(err)
	}

	if user.Status != users.StatusActive {
		switch user.Status {
		case users.StatusCreated:
			return "", ErrPermission.New("Users email not confirmed")
		case users.StatusSuspended:
			return "", ErrPermission.New("User suspended")
		}
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		return "", ErrUnauthenticated.Wrap(err)
	}

	claims := auth.Claims{
		UserID:    user.ID,
		Email:     user.Email,
		ExpiresAt: time.Now().UTC().Add(TokenExpirationTime),
	}

	token, err = service.signer.CreateToken(ctx, &claims)
	if err != nil {
		return "", Error.Wrap(err)
	}

	err = service.users.UpdateLastLogin(ctx, user.ID)
	return token, Error.Wrap(err)
}

// PreAuthToken authenticates User by credentials and returns pre auth token.
func (service *Service) PreAuthToken(ctx context.Context, email string) (token string, err error) {
	user, err := service.users.GetByEmail(ctx, email)
	if err != nil {
		return "", Error.Wrap(err)
	}

	claims := auth.Claims{
		UserID:    user.ID,
		Email:     user.Email,
		ExpiresAt: time.Now().UTC().Add(PreAuthTokenExpirationTime),
	}

	token, err = service.signer.CreateToken(ctx, &claims)
	if err != nil {
		return "", Error.Wrap(err)
	}

	return token, nil
}

// Authorize validates token from context and returns authorized Authorization.
func (service *Service) Authorize(ctx context.Context, tokenS string) (_ auth.Claims, err error) {
	token, err := auth.FromBase64URLString(tokenS)
	if err != nil {
		return auth.Claims{}, Error.Wrap(err)
	}

	claims, err := service.authenticate(token)
	if err != nil {
		return auth.Claims{}, ErrUnauthenticated.Wrap(err)
	}

	err = service.authorize(ctx, claims)
	if err != nil {
		return auth.Claims{}, ErrUnauthenticated.Wrap(err)
	}

	return *claims, nil
}

// authenticate validates token signature and returns authenticated *satelliteauth.Authorization.
func (service *Service) authenticate(token auth.Token) (_ *auth.Claims, err error) {
	signature := token.Signature

	err = service.signer.SignToken(&token)
	if err != nil {
		return nil, err
	}

	if subtle.ConstantTimeCompare(signature, token.Signature) != 1 {
		return nil, errs.New("incorrect signature")
	}

	claims, err := auth.FromJSON(token.Payload)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// authorize checks claims and returns authorized User.
func (service *Service) authorize(ctx context.Context, claims *auth.Claims) (err error) {
	if !claims.ExpiresAt.IsZero() && claims.ExpiresAt.Before(time.Now()) {
		return ErrUnauthenticated.New("token expiration time has expired")
	}

	user, err := service.users.GetByEmail(ctx, claims.Email)
	if err != nil {
		return ErrUnauthenticated.New("authorization failed. no user with email: %s", claims.Email)
	}

	if user.Status != users.StatusActive {
		switch user.Status {
		case users.StatusCreated:
			return ErrPermission.New("Users email not confirmed")
		case users.StatusSuspended:
			return ErrPermission.New("User suspended")
		}
	}

	return nil
}

// Register - registers a new user.
func (service *Service) Register(ctx context.Context, email, password, nickName, firstName, lastName string, wallet cryptoutils.Address) error {
	// check if the user email address already exists.
	_, err := service.users.GetByEmail(ctx, email)
	if err == nil {
		return ErrAddressAlreadyInUse.New("Email address is already in use.")
	}

	// check the password is valid.
	if !users.IsPasswordValid(password) {
		return Error.New("The password must contain at least one lowercase (a-z) letter, one uppercase (A-Z) letter, one digit (0-9) and one special character.")
	}

	user := users.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: []byte(password),
		NickName:     nickName,
		FirstName:    firstName,
		LastName:     lastName,
		LastLogin:    time.Time{},
		Status:       users.StatusCreated,
		CreatedAt:    time.Now().UTC(),
		// @TODO at the time of testing login through metamask.
		Wallet: wallet,
	}

	err = user.EncodePass()
	if err != nil {
		return Error.Wrap(err)
	}

	err = service.users.Create(ctx, user)
	if err != nil {
		return Error.Wrap(err)
	}

	token, err := service.Token(ctx, user.Email, password)
	if err != nil {
		return Error.Wrap(err)
	}

	// launch a goroutine that sends the email verification.
	go func() {
		err = service.emailService.SendVerificationEmail(user.Email, token)
		if err != nil {
			service.log.Error("Unable to send account activation email", Error.Wrap(err))
		}
	}()

	return err
}

// ConfirmUserEmail - parse token and confirm User.
func (service *Service) ConfirmUserEmail(ctx context.Context, activationToken string) error {
	token, err := auth.FromBase64URLString(activationToken)
	if err != nil {
		return Error.Wrap(err)
	}

	claims, err := service.authenticate(token)
	if err != nil {
		return ErrUnauthenticated.Wrap(err)
	}

	if !claims.ExpiresAt.IsZero() && claims.ExpiresAt.Before(time.Now()) {
		return ErrUnauthenticated.New("token expiration time has expired")
	}

	user, err := service.users.GetByEmail(ctx, claims.Email)
	if err != nil {
		return ErrUnauthenticated.Wrap(err)
	}

	if user.Status != users.StatusCreated {
		return ErrPermission.Wrap(err)
	}

	return Error.Wrap(service.users.Update(ctx, users.StatusActive, user.ID))
}

// ChangePassword - change users password.
func (service *Service) ChangePassword(ctx context.Context, password, newPassword string) error {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return ErrUnauthenticated.Wrap(err)
	}

	user, err := service.users.GetByEmail(ctx, claims.Email)
	if err != nil {
		return users.ErrUsers.Wrap(err)
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		return ErrUnauthenticated.Wrap(err)
	}

	user.PasswordHash = []byte(newPassword)
	err = user.EncodePass()
	if err != nil {
		return Error.Wrap(err)
	}

	return Error.Wrap(service.users.UpdatePassword(ctx, user.PasswordHash, user.ID))
}

// ResetPasswordSendEmail - send email with token for user.
func (service *Service) ResetPasswordSendEmail(ctx context.Context, email string) error {
	user, err := service.users.GetByEmail(ctx, email)
	if err != nil {
		return users.ErrUsers.Wrap(err)
	}

	if user.Status != users.StatusActive {
		switch user.Status {
		case users.StatusCreated:
			return ErrPermission.New("Users email not confirmed")
		case users.StatusSuspended:
			return ErrPermission.New("User suspended")
		}
	}

	token, err := service.PreAuthToken(ctx, user.Email)
	if err != nil {
		return Error.Wrap(err)
	}

	go func() {
		err = service.emailService.SendResetPasswordEmail(user.Email, token)
		if err != nil {
			service.log.Error("Unable to send reset password email", Error.Wrap(err))
		}
	}()

	return err
}

// CheckAuthToken checks auth token.
func (service *Service) CheckAuthToken(ctx context.Context, tokenStr string) error {
	token, err := auth.FromBase64URLString(tokenStr)
	if err != nil {
		return Error.Wrap(err)
	}
	claims, err := service.authenticate(token)
	if err != nil {
		return ErrUnauthenticated.Wrap(err)
	}

	if !claims.ExpiresAt.IsZero() && claims.ExpiresAt.Before(time.Now()) {
		return ErrUnauthenticated.New("token expiration time has expired")
	}

	_, err = service.users.GetByEmail(ctx, claims.Email)
	if err != nil {
		return users.ErrUsers.Wrap(err)
	}

	return nil
}

// ResetPassword - changes users password.
func (service *Service) ResetPassword(ctx context.Context, newPassword string) error {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return ErrUnauthenticated.Wrap(err)
	}

	user, err := service.users.GetByEmail(ctx, claims.Email)
	if err != nil {
		return users.ErrUsers.Wrap(err)
	}

	user.PasswordHash = []byte(newPassword)
	err = user.EncodePass()
	if err != nil {
		return Error.Wrap(err)
	}

	return Error.Wrap(service.users.UpdatePassword(ctx, user.PasswordHash, user.ID))
}

// TokenMessage creates message token and send to metamask for login.
func (service *Service) TokenMessage(ctx context.Context) (token string, err error) {
	claims := auth.Claims{
		ExpiresAt: time.Now().UTC().Add(PreAuthTokenExpirationTime),
	}

	token, err = service.signer.CreateToken(ctx, &claims)
	return token, Error.Wrap(err)
}

// CheckMetamaskTokenMessage - parses token-message and checks for expiration time.
func (service *Service) CheckMetamaskTokenMessage(ctx context.Context, tokenMessage string) error {
	token, err := auth.FromBase64URLString(tokenMessage)
	if err != nil {
		return Error.Wrap(err)
	}

	claims, err := service.authenticate(token)
	if err != nil {
		return ErrUnauthenticated.Wrap(err)
	}

	if !claims.ExpiresAt.IsZero() && claims.ExpiresAt.Before(time.Now()) {
		return ErrUnauthenticated.New("token expiration time has expired")
	}

	return Error.Wrap(err)
}

// LoginWithMetamask authenticates user by credentials and returns login token.
func (service *Service) LoginWithMetamask(ctx context.Context, loginMetamaskFields users.LoginMetamaskFields) (token string, err error) {
	verifyLoginMetamaskFields, err := verifyLoginMetamaskFields(loginMetamaskFields)
	if err != nil {
		return "", Error.Wrap(err)
	}

	if !verifyLoginMetamaskFields {
		return "", Error.New("login metamask fields are wrong")
	}

	wallet := cryptoutils.Address(strings.ToLower(string(loginMetamaskFields.Address)))

	user, err := service.users.GetByWalletAddress(ctx, loginMetamaskFields.Address)
	switch {
	case users.ErrNoUser.Has(err):
		user = users.User{
			ID:        uuid.New(),
			LastLogin: time.Time{},
			Status:    users.StatusActive,
			CreatedAt: time.Now().UTC(),
			Wallet:    wallet,
		}
		err = service.users.Create(ctx, user)
		if err != nil {
			return "", Error.Wrap(err)
		}
	case err != nil:
		return "", Error.Wrap(err)
	}

	claims := auth.Claims{
		UserID:    user.ID,
		ExpiresAt: time.Now().UTC().Add(TokenExpirationTime),
	}

	token, err = service.signer.CreateToken(ctx, &claims)
	if err != nil {
		return "", Error.Wrap(err)
	}

	err = service.users.UpdateLastLogin(ctx, user.ID)

	return token, Error.Wrap(err)
}

// verifyLoginMetamaskFields function that verifies the authenticity of the address.
func verifyLoginMetamaskFields(loginMetamaskFields users.LoginMetamaskFields) (bool, error) {
	fromAddr := common.HexToAddress(string(loginMetamaskFields.Address))
	hash := hexutil.MustDecode(loginMetamaskFields.Hash)

	if hash[64] != 27 && hash[64] != 28 {
		return false, Error.New("hash is wrong")
	}
	hash[64] -= 27

	pubKey, err := crypto.SigToPub(cryptoutils.SignHash([]byte(loginMetamaskFields.Message)), hash)
	if err != nil {
		return false, Error.Wrap(err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return fromAddr == recoveredAddr, nil
}
