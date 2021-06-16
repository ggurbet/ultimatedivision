// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"strings"
)

// normalizeEmail brings the email to UpperCase.
func normalizeEmail(email string) string {
	return strings.ToUpper(email)
}
