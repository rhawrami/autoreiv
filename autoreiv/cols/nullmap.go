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
	byteIdx, shiftBy := i/8, i%8
	ogVal, valIfNull := m[byteIdx], m[byteIdx]&^(1<<shiftBy)
	return ogVal == valIfNull
}

// SetNull sets a corresponding column record to null
func (m NullMap) SetNull(i int) {
	byteIdx, shiftBy := i/8, i%8
	m[byteIdx] = m[byteIdx] &^ (1 << shiftBy)
}

// SetNotNull sets a corresponding column record to not-null
func (m NullMap) SetNotNull(i int) {
	byteIdx, shiftBy := i/8, i%8
	m[byteIdx] = m[byteIdx] | (1 << shiftBy)
}

// NewNullMapFromBool returns a new NullMap, taking in a
// bool slice as input.
func NewNullMapFromBool(b []bool) NullMap {
	// more likely than not to not be div by 8
	// add remainder first, subtract if needed
	lenMap := len(b)/8 + 1
	if len(b)%8 == 0 {
		lenMap = lenMap - 1
	}

	m := make(NullMap, lenMap)
	for i := 0; i < len(b); i++ {
		if b[i] {
			bIdx, shiftBy := i/8, i%8
			// flip zero bit
			m[bIdx] = m[bIdx] | (1 << shiftBy)
		}
	}
	return m
}
