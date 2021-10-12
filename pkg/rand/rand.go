// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package rand

import (
	"fmt"
	"math/rand"
)

// RandomInRange randomizes numbers in the specified range.
func RandomInRange(count int) (int, error) {
	switch {
	case count == 1:
		return 1, nil
	case count > 0:
		return rand.Intn(count-1) + 1, nil
	default:
		return 0, fmt.Errorf("the number is less than or equal to zero")
	}
}

// IsIncludeRange indicates if probability includes percent.
func IsIncludeRange(percent int) bool {
	if (rand.Intn(99) + 1) <= percent {
		return true
	}
	return false
}
