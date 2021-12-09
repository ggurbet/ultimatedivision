// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package udts

import (
	"math/big"

	"github.com/google/uuid"
)

// UDT entity describes how many tokens of udt and what nonce the user has.
type UDT struct {
	UserID uuid.UUID `json:"userId"`
	Value  big.Int   `json:"value"`
	Nonce  int64     `json:"nonce"`
}
