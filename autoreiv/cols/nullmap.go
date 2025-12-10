package cols

// NullMap represents a bitmap of null values corresponding to
// a dataframe column.
//
// Taken after Apache Arrow's "validity map"
type NullMap []byte

// Len returns the byte-length (e.g., not "true" length) of the NullMap
func (m NullMap) Len() int {
	return len(m)
}

// LookUp returns whether a corresponding column record `i` is null.
//
// Assumes record at `i` exists
func (m NullMap) IsNull(i int) bool {
	// find which byte contains corresponding bit, and corresponding bit offset
	// e.g., find if column record at idx 10 is null:
	// 1. find that col idx 10 is in byte idx 1
	// 2. shift byte idx 1 by 2 units to get bit position 2
	// 3. return true if bit value is zero
	byteIdx, shiftBy := i/8, i%8
	return (m[byteIdx]>>shiftBy)&1 == 0
}

// SetNull sets a corresponding column record to null
//
// if already null, does nothing
func (m NullMap) SetNull(i int) {
	// find position of within-map byte and within-byte bit
	byteIdx, shiftBy := i/8, i%8
	m[byteIdx] = m[byteIdx] + (1 << shiftBy)
}

// SetNotNull sets a corresponding column record to not-null
func (m NullMap) SetNotNull(i int) {
	// find position of within-map byte and within-byte bit
	byteIdx, shiftBy := i/8, i%8
	m[byteIdx] = m[byteIdx] + (1 << shiftBy)
}
