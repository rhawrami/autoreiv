package main

import (
	"fmt"

	"github.com/rhawrami/autoreiv/readers"
)

func main() {
	sampleDat := []byte(`1001,Alice Johnson,28,75000.50,true,2020-03-15,Engineering,"Works on backend systems, excellent performer"
1002,Bob Smith,35,92000.00,true,2018-07-22,Sales,Lead sales representative
1003,"Chen, Wei",42,105000.75,true,2015-01-10,Management,"Director of Operations, ""exceptional"" leadership"
1004,Diana Prince,31,68000.00,false,2019-11-03,Marketing,Left for another opportunity
1005,Edward ""Ed"" Martinez,29,71500.25,true,2021-06-18,Engineering,Junior developer
1006,Fatima Hassan,38,88000.50,true,2017-09-05,Finance,Senior accountant
1007,George Taylor,45,125000.00,true,2012-04-30,Management,Chief Technology Officer
1008,Hannah Lee,26,62000.00,true,2022-02-14,Marketing,"Social media specialist
Handles multiple platforms"
`)
	colNames := readers.ParseRows(sampleDat, byte(','), string("\n"))
	for _, col := range colNames {
		fmt.Printf("%#v\n", col)
		fmt.Println(len(col))
	}
}
