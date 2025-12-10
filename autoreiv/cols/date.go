package cols

import (
	"fmt"
	"time"
)

// TimeCol represents a column of Time values, using the time.Time
// type internally
//
// Follows from Apache Arrow columnar format, with separate
// value and null bit-mask arrays
type TimeCol struct {
	Vals  []time.Time // actual values
	Nulls []bool      // null bit-map
}

// Len returns the record length of a column
func (c *TimeCol) Len() int {
	return len(c.Vals)
}

// GetVals returns a column's raw values
func (c *TimeCol) GetVals() []time.Time {
	return c.Vals
}

// GetNulls returns a column's null bit-map
func (c *TimeCol) GetNulls() []bool {
	return c.Nulls
}

// IsNull returns if a given value is null
//
// returns error if index is out-of-bounds or
// if index is negative
func (c *TimeCol) IsNull(i int) (bool, error) {
	if i >= len(c.Vals) || i < 0 {
		return false, fmt.Errorf("index %d not in i-range [0,%d]", i, len(c.Vals)-1)
	}
	return c.Nulls[i], nil
}
