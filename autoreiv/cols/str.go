package cols

import (
	"fmt"
)

// StrCol represents a column of string values
//
// Follows from Apache Arrow columnar format, with separate
// value and null bit-mask arrays
type StrCol struct {
	Vals  []string // actual values
	Nulls []bool   // null bit-map
}

// Len returns the record length of a column
func (c *StrCol) Len() int {
	return len(c.Vals)
}

// GetVals returns a column's raw values
func (c *StrCol) GetVals() []string {
	return c.Vals
}

// GetNulls returns a column's null bit-map
func (c *StrCol) GetNulls() []bool {
	return c.Nulls
}

// IsNull returns if a given value is null
//
// returns error if index is out-of-bounds or
// if index is negative
func (c *StrCol) IsNull(i int) (bool, error) {
	if i >= len(c.Vals) || i < 0 {
		return false, fmt.Errorf("index %d not in i-range [0,%d]", i, len(c.Vals)-1)
	}
	return c.Nulls[i], nil
}
