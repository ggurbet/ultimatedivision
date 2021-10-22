// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package mail

// ensures that SMTPSender implements Sender.
var _ Sender = (*MockSender)(nil)

// MockSender is mock implementation of Sender.
type MockSender struct{}

// FromAddress returns address of sender from.
func (sender *MockSender) FromAddress() Address {
	return Address{
		Name:    "",
		Address: "",
	}
}

// SendEmail sends email message to the given recipient.
func (sender *MockSender) SendEmail(msg *Message) error {
	return nil
}
