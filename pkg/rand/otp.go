// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package rand

import (
	"crypto/rand"
	"io"

	"github.com/zeebo/errs"
)

var (
	// Error indicates about OTP generation error.
	Error = errs.Class("OTP generation failed")
	// ValidationError indicates that OTP is not valid.
	ValidationError = errs.Class("OTP is not valid")
)

const otpLength = 6

// OTP is a 6 random digits.
type OTP string

// NewOTP generates otp - 6 random string.
func NewOTP() (OTP, error) {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	b := make([]byte, otpLength)

	n, err := io.ReadAtLeast(rand.Reader, b, otpLength)
	if n != otpLength {
		return "", Error.Wrap(err)
	}

	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return OTP(b), nil
}

// ValidateOTP could be used to validate OTP code.
func ValidateOTP(otp OTP) error {
	if len(otp) != otpLength {
		return ValidationError.New("wrong length")
	}

	for _, c := range otp {
		if c < '0' || c > '9' {
			return ValidationError.New("contains wrong symbols")
		}
	}

	return nil
}
