package readers

import "strconv"

type ARType int

const (
	ARInteger ARType = iota
	ARString
	ARFloat
	ARBool
	ARDate
)

func ConvFloat(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}
