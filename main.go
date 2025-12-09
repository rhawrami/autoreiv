package main

import (
	"fmt"

	"github.com/rhawrami/autoreiv/readers"
)

func main() {
	datesToTest := []string{"2010-27-11", "09/20/2020", "1983-30-01"}
	for _, s := range datesToTest {
		d, err := readers.NewdDate(s)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			fmt.Printf("%s -> %#v\n", s, d)
		}
	}
}
