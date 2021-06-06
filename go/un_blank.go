package demo

import "fmt"

func DeBlank() {
	ars := Int{1, 2, 3, 4}
	brs := Int{1, 2, 4, 8}

	for i, _ := range ars {
		if !brs.Present(ars[i]) {
			fmt.Printf("%d from %v is not present in %v\n", ars[i], ars, brs)
		}
	}
}
