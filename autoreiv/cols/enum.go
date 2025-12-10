package cols

import (
	"fmt"
)

// EnumCol represents a column of enumerated values
//
// the Dict includes all possible values, and Vals include what
// value (by Dict index lookup) a record actually has
//
// Follows from Apache Arrow columnar format, with separate
// value and null bit-mask arrays
type EnumCol struct {
	Dict  []string // possible values
	Vals  []int    // actual value (from Dict idx)
	Nulls []bool   // null bit-map
}

// Len returns the record length of a column
func (c *EnumCol) Len() int {
	return len(c.Vals)
}

// GetVals returns a column's raw values
func (c *EnumCol) GetVals() []int {
	return c.Vals
}

// GetNulls returns a column's null bit-map
func (c *EnumCol) GetNulls() []bool {
	return c.Nulls
}

// IsNull returns if a given value is null
//
// returns error if index is out-of-bounds or
// if index is negative
func (c *EnumCol) IsNull(i int) (bool, error) {
	if i >= len(c.Vals) || i < 0 {
		return false, fmt.Errorf("index %d not in i-range [0,%d]", i, len(c.Vals)-1)
	}
	return c.Nulls[i], nil
}
