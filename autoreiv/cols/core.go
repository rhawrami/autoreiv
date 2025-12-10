package cols

// Col represents a dataframe column
type Col interface {
	Len() int
	IsNull(i int) (bool, error)
}
