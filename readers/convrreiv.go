package readers

import (
	"fmt"
	"regexp"
	"strconv"
)

// dType specifies the currently supported dataframe types:
// {integer, string, float, boolean, date}
// There is also the dNaN for null values
type dType int

const (
	dNaN dType = iota
	dInt64
	dString
	dFloat64
	dBool
)

// stringToInt64 is a convenience function for...well it's in the name
func stringToInt64(s string) (int64, error) {
	// assume base10
	base := 10
	if x, err := strconv.ParseInt(s, base, 64); err == nil {
		return x, nil
	} else {
		return 0, err
	}
}

// stringToFloat64 is a convenience function for...stringToFloat64 is a convenience
// function for...stringToFloat64 is a convenience function for...
func stringToFloat64(s string) (float64, error) {
	if x, err := stringToFloat64(s); err == nil {
		return x, nil
	} else {
		return 0, err
	}
}

// come back to this
func stringToBool(s string) (bool, error) {
	if isTrue, err := regexp.Match(`[Tt]`, []byte(s)); isTrue {
		return true, nil
	} else if !isTrue && err == nil {
		return false, nil
	} else {
		return false, err
	}
}

type dDate struct {
	year  int64
	month int64
	day   int64
}

func NewdDate(s string) (dDate, error) {
	shortM, _ := regexp.Compile(`(\d{1,4})[-\/](\d{1,2})[-\/](\d{1,4})`)
	shortMatches := shortM.FindStringSubmatch(s)
	if shortMatches == nil {
		return dDate{}, fmt.Errorf("string does not match date format: %s", s)
	}
	// must have len 4 if perfectly matched
	if len(shortMatches) != 4 {
		return dDate{}, fmt.Errorf("string does not match date format: %s", s)
	}

	// assume MM DD YYYY format
	sMonth, sDay, sYear := shortMatches[1], shortMatches[2], shortMatches[3]
	// if first grp match has >2 chars, then YYYY MM DD
	if len(shortMatches[1]) > 2 {
		sYear, sMonth, sDay = shortMatches[1], shortMatches[2], shortMatches[3]
	}
	var (
		mErr error
		dErr error
		yErr error
	)
	iMonth, mErr := (stringToInt64(sMonth))
	iDay, dErr := stringToInt64(sDay)
	iYear, yErr := stringToInt64(sYear)
	if mErr != nil || dErr != nil || yErr != nil {
		return dDate{}, fmt.Errorf("unable to convert string to date: %s", s)
	}

	// final basic test: if iMonth > 12 -> must be day -> swap iDay and iMonth
	if iMonth > 12 {
		iDay, iMonth = iMonth, iDay
	}

	return dDate{year: iYear, month: iMonth, day: iDay}, nil
}
