// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package bigint

import (
	"math/big"
)

// Max returns maximal element of slice.
func Max(elements []big.Int) big.Int {
	var max big.Int

	for _, value := range elements {
		if value.Cmp(&max) == 1 {
			max = value
		}
	}

	return max
}
