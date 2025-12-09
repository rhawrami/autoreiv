package readers

import (
	"regexp"
)

type arType int

const (
	arInteger arType = iota
	arString
	arFloat
	arBool
	arDate
)

func InferType(val []byte) arType {
	// float type
	if floatMatch, _ := regexp.Match(`^[\+-]?[0-9]+?[,0-9]+\.`, val); floatMatch {
		return arFloat
	}
	// int type
	if intMatch, _ := regexp.Match(`^[\+-]\d+,?`, val); intMatch {
		return arInteger
	}
	// bool type (deal with case sensitive later)
	if boolMatch, _ := regexp.Match(`^(true|false)$`, val); boolMatch {
		return arBool
	}
	// date type (rudimentary for now)
	dateMatchShort, _ := regexp.Match(`\d{1,4}[-\/]\d{1,4}[-\/]\d{1,4}`, val)
	dateMatchLong, _ := regexp.Match(`(jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec)[a-z]+ \d{1,2},?\s\d{1,4}`, val)
	if dateMatchShort || dateMatchLong {
		return arDate
	}
	// string type as default
	return arString
}

type InfererReiv struct {
	colNames   []string
	sampleData []byte
	delimiter  string
}

// func NewInfererReiv(fileName string, delimiter, newlinechar byte) (InfererReiv, error) {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return InfererReiv{}, err
// 	}

// 	// read 4 KiB
// 	bytesToRead := 4000
// 	buffer := make([]byte, bytesToRead)
// 	n, err := file.Read(buffer)
// 	if err != nil {
// 		if err != io.EOF {
// 			return InfererReiv{}, err
// 		}
// 	}

// 	// get colNames

// }

func ParseRows(sample []byte, delimiter byte, newlinechars string) [][]string {
	rows := make([][]string, 0)
	// track multi-word entries
	needClosingQuote := false
	// track quoted words
	entryIsQuoted := false
	// start of entry
	startPos := 0

	row := make([]string, 0, 50)
	for i, v := range sample {
		// move to next line
		if string(v) == newlinechars && !needClosingQuote {
			// add final entry

			var entry []byte
			if sample[i-1] == '"' {
				entry = sample[startPos+1 : i-1]
			} else {
				entry = sample[startPos:i]
			}
			row = append(row, string(entry))

			// append row to rows, clear out row
			rows = append(rows, row)
			row = make([]string, 0, 50)

			startPos = i + 1
			continue
		}
		// read quotation
		if v == '"' {
			if !needClosingQuote {
				needClosingQuote = true
			} else {
				entryIsQuoted = true
				needClosingQuote = false
			}
		}
		// hit delimiter, add entry
		if v == delimiter && !needClosingQuote {
			var entry []byte
			if entryIsQuoted {
				entry = sample[startPos+1 : i-1]
			} else {
				entry = sample[startPos:i]
			}
			row = append(row, string(entry))
			// update state
			entryIsQuoted = false
			startPos = i + 1
		}
	}
	// if there are any remaining rows
	if len(row) > 0 {
		rows = append(rows, row)
	}
	return rows
}

func ParseRow(sample []byte, delimiter byte) []string {
	records := make([]string, 0)
	// track multi-word entries
	needClosingQuote := false
	// track quoted words
	entryIsQuoted := false
	// start of entry
	startPos := 0

	for i, v := range sample {
		if v == '"' {
			if !needClosingQuote {
				needClosingQuote = true
			} else {
				entryIsQuoted = true
				needClosingQuote = false
			}
		}
		// hit delimiter, add entry
		if v == delimiter && !needClosingQuote {
			var entry []byte
			if entryIsQuoted {
				entry = sample[startPos+1 : i-1]
			} else {
				entry = sample[startPos:i]
			}
			records = append(records, string(entry))
			// update state
			entryIsQuoted = false
			startPos = i + 1
		}
	}
	// final entry
	if sample[len(sample)-1] != delimiter {
		var entry []byte
		// if entry is quoted
		if sample[len(sample)-1] == '"' {
			entry = sample[startPos+1 : len(sample)-1]
		} else {
			entry = sample[startPos:]
		}
		records = append(records, string(entry))
	}
	return records
}
