// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package userauth

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/subtle"
	"strings"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/zeebo/errs"
	"golang.org/x/crypto/bcrypt"

	"ultimatedivision/console/emails"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
	"ultimatedivision/pkg/publicprivatekey"
	"ultimatedivision/pkg/velas"
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
// architecture: Service.
type Service struct {
	users        users.DB
	signer       auth.TokenSigner
	emailService *emails.Service
	log          logger.Logger
	velas        *velas.Service
}

// NewService is a constructor for user auth service.
func NewService(users users.DB, signer auth.TokenSigner, emails *emails.Service, log logger.Logger, velas *velas.Service) *Service {
	return &Service{
		users:        users,
		signer:       signer,
		emailService: emails,
		log:          log,
		velas:        velas,
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
func (service *Service) Register(ctx context.Context, email, password, nickName, firstName, lastName string, wallet common.Address) error {
	// check if the user email address already exists.
	_, err := service.users.GetByEmail(ctx, email)
	if err == nil {
		return ErrAddressAlreadyInUse.New("email address is already in use.")
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
		Wallet:     wallet,
		WalletType: users.WalletTypeETH,
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

// Nonce creates nonce and send to metamask for login.
func (service *Service) Nonce(ctx context.Context, address common.Address, walletType users.WalletType) (string, error) {
	user, err := service.users.GetByWalletAddress(ctx, address.String(), walletType)
	if err != nil {
		return "", Error.Wrap(err)
	}

	nonce := hexutil.Encode(user.Nonce)

	return nonce, nil
}

// RegisterWithMetamask creates user by credentials.
func (service *Service) RegisterWithMetamask(ctx context.Context, signature []byte) error {
	walletAddress, err := recoverWalletAddress([]byte(users.DefaultMessageForRegistration), signature)
	if err != nil {
		return Error.Wrap(err)
	}

	_, err = service.users.GetByWalletAddress(ctx, walletAddress.String(), users.WalletTypeETH)
	if !users.ErrNoUser.Has(err) {
		return Error.New("this user already exist")
	}

	nonce := make([]byte, 32)
	_, err = rand.Read(nonce)
	if err != nil {
		return Error.Wrap(err)
	}

	user := users.User{
		ID:         uuid.New(),
		Nonce:      nonce,
		LastLogin:  time.Time{},
		Status:     users.StatusActive,
		CreatedAt:  time.Now().UTC(),
		Wallet:     walletAddress,
		WalletType: users.WalletTypeETH,
	}
	err = service.users.Create(ctx, user)
	if err != nil {
		return Error.Wrap(err)
	}

	return nil
}

// LoginWithMetamask authenticates user by credentials and returns login token.
func (service *Service) LoginWithMetamask(ctx context.Context, nonce string, signature []byte) (string, error) {
	walletAddress, err := recoverWalletAddress([]byte(nonce), signature)
	if err != nil {
		return "", Error.Wrap(err)
	}

	user, err := service.users.GetByWalletAddress(ctx, walletAddress.String(), users.WalletTypeETH)
	if err != nil {
		return "", Error.Wrap(err)
	}

	decodeNonce, err := hexutil.Decode(nonce)
	if err != nil {
		return "", Error.Wrap(err)
	}

	if !bytes.Equal(decodeNonce, user.Nonce) {
		return "", Error.New("nonce is invalid")
	}

	claims := auth.Claims{
		UserID:    user.ID,
		ExpiresAt: time.Now().UTC().Add(TokenExpirationTime),
	}

	token, err := service.signer.CreateToken(ctx, &claims)
	if err != nil {
		return "", Error.Wrap(err)
	}

	newNonce := make([]byte, 32)
	_, err = rand.Read(newNonce)
	if err != nil {
		return "", Error.Wrap(err)
	}

	err = service.users.UpdateNonce(ctx, user.ID, newNonce)
	if err != nil {
		return "", Error.Wrap(err)
	}

	err = service.users.UpdateLastLogin(ctx, user.ID)
	if err != nil {
		service.log.Error("could not update last login", Error.Wrap(err))
	}

	return token, nil
}

// recoverWalletAddress function that verifies the authenticity of the address.
func recoverWalletAddress(message, signature []byte) (common.Address, error) {
	if signature[64] != 27 && signature[64] != 28 {
		return common.Address{}, Error.New("hash is wrong")
	}
	signature[64] -= 27

	pubKey, err := crypto.SigToPub(evmsignature.SignHash(message), signature)
	if err != nil {
		return common.Address{}, Error.Wrap(err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return recoveredAddr, nil
}

// SendEmailForChangeEmail - sends email for change users email address.
func (service *Service) SendEmailForChangeEmail(ctx context.Context, newEmail string) error {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return ErrUnauthenticated.Wrap(err)
	}

	// check if the new user email address already exists.
	_, err = service.users.GetByEmail(ctx, newEmail)
	if err == nil {
		return ErrAddressAlreadyInUse.New("email address is already in use.")
	}

	user, err := service.users.GetByEmail(ctx, claims.Email)
	if err != nil {
		return users.ErrUsers.Wrap(err)
	}

	token, err := service.PreAuthTokenToChangeEmail(ctx, user.Email, newEmail)
	if err != nil {
		return Error.Wrap(err)
	}

	go func() {
		err = service.emailService.SendVerificationEmailForChangeEmail(newEmail, token)
		if err != nil {
			service.log.Error("Unable to send verification email", Error.Wrap(err))
		}
	}()

	return Error.Wrap(err)
}

// ChangeEmail - changes users email address.
func (service *Service) ChangeEmail(ctx context.Context, activationToken string) error {
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

	user, err := service.users.Get(ctx, claims.UserID)
	if err != nil {
		return ErrUnauthenticated.Wrap(err)
	}

	return Error.Wrap(service.users.UpdateEmail(ctx, user.ID, claims.Email))
}

// PreAuthTokenToChangeEmail authenticates User by credentials and returns pre auth token with new email address.
func (service *Service) PreAuthTokenToChangeEmail(ctx context.Context, email, newEmail string) (token string, err error) {
	user, err := service.users.GetByEmail(ctx, email)
	if err != nil {
		return "", Error.Wrap(err)
	}

	claims := auth.Claims{
		UserID:    user.ID,
		Email:     newEmail,
		ExpiresAt: time.Now().Add(PreAuthTokenExpirationTime),
	}

	token, err = service.signer.CreateToken(ctx, &claims)
	if err != nil {
		return "", Error.Wrap(err)
	}

	return token, nil
}

// RegisterWithVelas creates user by credentials.
func (service *Service) RegisterWithVelas(ctx context.Context, walletAddress common.Address) error {
	_, err := service.users.GetByWalletAddress(ctx, walletAddress.String(), users.WalletTypeVelas)
	if !users.ErrNoUser.Has(err) {
		return Error.New("this user already exist")
	}

	nonce := make([]byte, 32)
	_, err = rand.Read(nonce)
	if err != nil {
		return Error.Wrap(err)
	}

	user := users.User{
		ID:         uuid.New(),
		Nonce:      nonce,
		LastLogin:  time.Time{},
		Status:     users.StatusActive,
		CreatedAt:  time.Now().UTC(),
		Wallet:     walletAddress,
		WalletType: users.WalletTypeVelas,
	}
	err = service.users.Create(ctx, user)
	if err != nil {
		return Error.Wrap(err)
	}

	return nil
}

// SaveVelasData save json to db while register velas user.
func (service *Service) SaveVelasData(ctx context.Context, walletAddress common.Address, velasString string) error {
	user, err := service.users.GetByWalletAddress(ctx, walletAddress.String(), users.WalletTypeVelas)
	if err != nil {
		return Error.New("can't get user by wallet")
	}

	var velasData = users.VelasData{
		ID:       user.ID,
		Response: velasString,
	}

	err = service.users.SetVelasData(ctx, velasData)
	if err != nil {
		return Error.Wrap(err)
	}

	return nil
}

// GetVelasData returns velas data by userId.
func (service *Service) GetVelasData(ctx context.Context, userID uuid.UUID) (users.VelasData, error) {
	velasData, err := service.users.GetVelasData(ctx, userID)
	return velasData, Error.Wrap(err)
}

// LoginWithVelas authenticates user by credentials and returns login token.
func (service *Service) LoginWithVelas(ctx context.Context, nonce string, walletAddress common.Address) (string, error) {
	user, err := service.users.GetByWalletAddress(ctx, walletAddress.String(), users.WalletTypeVelas)
	if err != nil {
		return "", Error.Wrap(err)
	}

	decodeNonce, err := hexutil.Decode(nonce)
	if err != nil {
		return "", Error.Wrap(err)
	}

	if !bytes.Equal(decodeNonce, user.Nonce) {
		return "", Error.New("nonce is invalid")
	}

	claims := auth.Claims{
		UserID:    user.ID,
		ExpiresAt: time.Now().UTC().Add(TokenExpirationTime),
	}

	token, err := service.signer.CreateToken(ctx, &claims)
	if err != nil {
		return "", Error.Wrap(err)
	}

	newNonce := make([]byte, 32)
	_, err = rand.Read(newNonce)
	if err != nil {
		return "", Error.Wrap(err)
	}

	err = service.users.UpdateNonce(ctx, user.ID, newNonce)
	if err != nil {
		return "", Error.Wrap(err)
	}

	err = service.users.UpdateLastLogin(ctx, user.ID)
	if err != nil {
		service.log.Error("could not update last login", Error.Wrap(err))
	}

	return token, nil
}

// VelasVAClientFields returns velas va client fields.
func (service *Service) VelasVAClientFields() velas.VAClientFields {
	return service.velas.Get()
}

// RegisterWithCasper creates user by credentials.
func (service *Service) RegisterWithCasper(ctx context.Context, walletAddress string) error {
	_, err := service.users.GetByWalletAddress(ctx, walletAddress, users.WalletTypeCasper)
	if !users.ErrNoUser.Has(err) {
		return Error.New("this user already exist")
	}

	publicKey, privateKey, err := publicprivatekey.GeneratePublicPrivateKey()
	if err != nil {
		return Error.Wrap(err)
	}

	user := users.User{
		ID:             uuid.New(),
		PublicKey:      string(publicKey),
		PrivateKey:     string(privateKey),
		LastLogin:      time.Time{},
		Status:         users.StatusActive,
		CreatedAt:      time.Now().UTC(),
		CasperWallet:   walletAddress,
		CasperWalletID: walletAddress,
		WalletType:     users.WalletTypeCasper,
	}
	err = service.users.Create(ctx, user)
	if err != nil {
		return Error.Wrap(err)
	}

	return nil
}

// LoginWithCasper authenticates user by credentials and returns login token.
func (service *Service) LoginWithCasper(ctx context.Context, publicKey string, signature string) (string, error) {
	key, err := service.users.GetByPublicKey(ctx, publicKey)
	if err != nil {
		return "", Error.Wrap(err)
	}

	walletAddress, err := publicprivatekey.DecryptCasperWalletAddress(signature, []byte(key.PrivateKey))
	if err != nil {
		return "", Error.New("invalid signature")
	}

	user, err := service.users.GetByWalletAddress(ctx, string(walletAddress), users.WalletTypeCasper)
	if err != nil {
		return "", Error.Wrap(err)
	}

	if !strings.EqualFold(key.PrivateKey, user.PrivateKey) {
		return "", Error.New("nonce is invalid")
	}

	claims := auth.Claims{
		UserID:    user.ID,
		ExpiresAt: time.Now().UTC().Add(TokenExpirationTime),
	}

	token, err := service.signer.CreateToken(ctx, &claims)
	if err != nil {
		return "", Error.Wrap(err)
	}

	newPublicKey, newPrivateKey, err := publicprivatekey.GeneratePublicPrivateKey()
	if err != nil {
		return "", Error.Wrap(err)
	}

	err = service.users.UpdatePublicPrivateKey(ctx, user.ID, string(newPublicKey), string(newPrivateKey))
	if err != nil {
		return "", Error.Wrap(err)
	}

	err = service.users.UpdateLastLogin(ctx, user.ID)
	if err != nil {
		service.log.Error("could not update last login", Error.Wrap(err))
	}

	return token, nil
}

// PublicKey get public key and send to casper for login.
func (service *Service) PublicKey(ctx context.Context, address string) (string, error) {
	user, err := service.users.GetByWalletAddress(ctx, address, users.WalletTypeCasper)
	if err != nil {
		return "", Error.Wrap(err)
	}

	return user.PublicKey, nil
}
