package autoreiv

type AutoReiv[T any] struct {
	schema schematic
	df     dataFrame[T]
}

type schematic struct {
	colPositions map[string][]recordPosition
}

type recordPosition struct {
	start int
	end   int
}

type dataFrame[T any] struct {
	data map[string]column[T]
}

type column[T any] struct {
	vals []T
}
