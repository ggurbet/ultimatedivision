// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package mail

import (
	"strings"
)

// Normalize brings the email to UpperCase.
func Normalize(email string) string {
	return strings.ToUpper(email)
}
