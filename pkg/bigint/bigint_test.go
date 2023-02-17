// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package bigint

import (
	"math/big"
	"testing"
)

func TestMax(t *testing.T) {
	type testpair struct {
		dataset []big.Int
		res     big.Int
	}

	big1 := big.NewInt(1500)
	big2 := big.NewInt(1200)
	big3 := big.NewInt(1300)

	testcases := []testpair{
		{
			dataset: []big.Int{*big1, *big2, *big3},
			res:     *big1,
		},
		{
			dataset: nil,
			res:     big.Int{},
		},
	}

	for _, testcase := range testcases {
		actualResult := Max(testcase.dataset)
		if actualResult.Cmp(&testcase.res) != 0 {
			t.Error(
				"For", testcase.dataset,
				"expected", testcase.res,
				"got", actualResult,
			)
		}
	}
}
