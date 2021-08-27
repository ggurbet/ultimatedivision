// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package sqlsearchoperators

// SearchOperators entity to mapping a list of possible search operators for sql.
var SearchOperators = map[string]SearchOperator{
	"eq":  EQ,
	"gt":  GT,
	"lt":  LT,
	"gte": GTE,
	"lte": LTE,
}

// SearchOperator defines the list of possible search operators for sql.
type SearchOperator string

const (
	// EQ - equal to value.
	EQ SearchOperator = "="
	// GT - greater than or equal to value.
	GT SearchOperator = ">"
	// LT - less than to value.
	LT SearchOperator = "<"
	// GTE - greater than  to value.
	GTE SearchOperator = ">="
	// LTE - less than or equal to value.
	LTE SearchOperator = "<="
	// LIKE - like to value.
	LIKE SearchOperator = "LIKE"
)
