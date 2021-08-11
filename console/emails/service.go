// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package emails

import (
	"fmt"
	"time"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/internal/mail"
)

// Config defines values needed by mailservice service.
// TODO: separate on service and client configs.
type Config struct {
	Provider          string `json:"provider" default:"mock"`
	SMTPServerAddress string `json:"smtpServerAddress" help:"smtp server address"`
	TemplatePath      string `json:"templatePath" help:"path to email templates source"`
	From              string `json:"from" help:"sender email address"`
	AuthType          string `json:"authType" help:"smtp authentication type"`
	PlainLogin        string `json:"plainLogin" help:"plain authentication user login"`
	PlainPassword     string `json:"plainPassword" help:"plain authentication user password"`
	RefreshToken      string `json:"refreshToken" help:"refresh token used to retrieve new access token"`
	ClientID          string `json:"clientId" help:"oauth2 app's client id"`
	ClientSecret      string `json:"clientSecret" help:"oauth2 app's client secret"`
	TokenURI          string `json:"tokenURI" help:"uri which is used when retrieving new access token"`

	TransactionsFileName string `json:"transactionsFileName" help:"name of file that will be attached to email"`
	Domain               string `json:"domain"`
}

// Error indicates about email sending error.
var Error = errs.Class("email service error")

// Service contains all business related logic.
//
// architecture: service
type Service struct {
	log    logger.Logger
	sender mail.Sender
	config Config
}

// NewService is a constructor for service.
func NewService(log logger.Logger, sender mail.Sender, config Config) *Service {
	return &Service{
		log:    log,
		sender: sender,
		config: config,
	}
}

// SendVerificationEmail is used to send email with verification link.
func (service *Service) SendVerificationEmail(email, token string) error {
	var verificationMessage mail.Message

	verificationMessage.To = []mail.Address{{Address: email, Name: "Verify"}}
	verificationMessage.Date = time.Now().UTC()
	verificationMessage.PlainText = fmt.Sprintf("%s/%s", service.config.Domain, token)
	verificationMessage.Subject = "confirm your email"
	verificationMessage.From = mail.Address{Address: service.config.From}

	return service.sender.SendEmail(&verificationMessage)
}
