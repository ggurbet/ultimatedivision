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
	SMTPServerAddress string `help:"smtp server address" default:""`
	TemplatePath      string `help:"path to email templates source" default:""`
	From              string `help:"sender email address" default:""`
	AuthType          string `help:"smtp authentication type" default:"simulate"`
	PlainLogin        string `help:"plain authentication user login" default:""`
	PlainPassword     string `help:"plain authentication user password" default:""`
	RefreshToken      string `help:"refresh token used to retrieve new access token" default:""`
	ClientID          string `help:"oauth2 app's client id" default:""`
	ClientSecret      string `help:"oauth2 app's client secret" default:""`
	TokenURI          string `help:"uri which is used when retrieving new access token" default:""`

	TransactionsFileName string `help:"name of file that will be attached to email" default:"transactions"`
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
	config *Config
}

// NewService is a constructor for service.
func NewService(log logger.Logger, sender mail.Sender, config *Config) *Service {
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
