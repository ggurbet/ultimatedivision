// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package userauth

import (
	"context"
	"crypto/subtle"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
	"golang.org/x/crypto/bcrypt"

	"ultimatedivision/console/emails"
	"ultimatedivision/internal/auth"
	"ultimatedivision/internal/logger"
	"ultimatedivision/users"
)

const (
	// TokenExpirationTime after passing this time token expires.
	TokenExpirationTime = 24 * time.Hour
)

var (
	// ErrUnauthenticated should be returned when user performs unauthenticated action.
	ErrUnauthenticated = errs.Class("user unauthenticated error")

	// Error is a error class for internal auth errors.
	Error = errs.Class("user auth internal error")

	// ErrPermission should be returned when user permission denied.
	ErrPermission = errs.Class("permission denied")
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
		Email:     user.Email,
		ExpiresAt: time.Now().Add(TokenExpirationTime),
	}

	token, err = service.signer.CreateToken(ctx, &claims)
	if err != nil {
		return "", Error.Wrap(err)
	}

	return token, nil
}

// Authorize validates token from context and returns authorized Authorization.
func (service *Service) Authorize(ctx context.Context) (_ auth.Claims, err error) {
	tokenS, ok := auth.GetToken(ctx)
	if !ok {
		return auth.Claims{}, ErrUnauthenticated.Wrap(err)
	}

	token, err := auth.FromBase64URLString(string(tokenS))
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
		return ErrUnauthenticated.Wrap(err)
	}

	_, err = service.users.GetByEmail(ctx, claims.Email)
	if err != nil {
		return errs.New("authorization failed. no user with email: %s", claims.Email)
	}

	return nil
}

// Register - register a new user.
func (service *Service) Register(ctx context.Context, email, password, nickName, firstName, lastName string) error {
	// check if the user email address already exists.
	_, err := service.users.GetByEmail(ctx, email)
	if err == nil {
		return Error.Wrap(errs.New("This email address is already in use."))
	}

	// check the password is valid.
	if !users.IsPasswordValid(password) {
		return Error.Wrap(errs.New("The password must contain at least one lowercase (a-z) letter, one uppercase (A-Z) letter, one digit (0-9) and one special character."))
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
	}

	err = user.EncodePass()
	if err != nil {
		return Error.Wrap(err)
	}

	err = service.users.Create(ctx, user)
	if err != nil {
		return Error.Wrap(err)
	}

	token, err := service.Token(ctx, user.Email, string(user.PasswordHash))
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
	status := int(users.StatusActive)
	token, err := auth.FromBase64URLString(activationToken)
	if err != nil {
		return Error.Wrap(err)
	}

	claims, err := service.authenticate(token)
	if err != nil {
		return ErrUnauthenticated.Wrap(err)
	}

	if !claims.ExpiresAt.IsZero() && claims.ExpiresAt.Before(time.Now()) {
		return ErrUnauthenticated.Wrap(err)
	}

	user, err := service.users.GetByEmail(ctx, claims.Email)
	if err != nil {
		return ErrUnauthenticated.Wrap(err)
	}

	if user.Status != users.StatusCreated {
		return ErrPermission.Wrap(err)
	}

	return Error.Wrap(service.users.Update(ctx, status, user.ID))
}
