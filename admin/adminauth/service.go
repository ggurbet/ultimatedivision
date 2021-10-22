// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package adminauth

import (
	"context"
	"crypto/subtle"
	"time"

	"github.com/zeebo/errs"
	"golang.org/x/crypto/bcrypt"

	"ultimatedivision/admin/admins"
	"ultimatedivision/pkg/auth"
)

const (
	// TokenExpirationTime after passing this time token expires.
	TokenExpirationTime = 24 * time.Hour
)

var (
	// ErrUnauthenticated should be returned when admin performs unauthenticated action.
	ErrUnauthenticated = errs.Class("admin unauthenticated error")

	// Error is a error class for internal auth errors.
	Error = errs.Class("admin auth internal error")
)

// Service is handling all admin authentication logic.
//
// architecture: Service
type Service struct {
	admins admins.DB
	signer auth.TokenSigner
}

// NewService is a constructor for admin auth service.
func NewService(admins admins.DB, signer auth.TokenSigner) *Service {
	return &Service{
		admins: admins,
		signer: signer,
	}
}

// Token authenticates Admin by credentials and returns auth token.
func (service *Service) Token(ctx context.Context, email string, password string) (token string, err error) {
	admin, err := service.admins.GetByEmail(ctx, email)
	if err != nil {
		return "", Error.Wrap(err)
	}

	err = bcrypt.CompareHashAndPassword(admin.PasswordHash, []byte(password))
	if err != nil {
		return "", ErrUnauthenticated.Wrap(err)
	}

	claims := auth.Claims{
		UserID:    admin.ID,
		Email:     admin.Email,
		ExpiresAt: time.Now().Add(TokenExpirationTime),
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
		return ErrUnauthenticated.Wrap(err)
	}

	_, err = service.admins.GetByEmail(ctx, claims.Email)
	if err != nil {
		return errs.New("authorization failed. no user with email: %s", claims.Email)
	}

	return nil
}
