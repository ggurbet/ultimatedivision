// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package mail

import (
	"net/mail"
)

// Address is alias of net/mail.Address.
type Address = mail.Address

// Sender sends emails.
type Sender interface {
	FromAddress() Address
	SendEmail(msg *Message) error
}
