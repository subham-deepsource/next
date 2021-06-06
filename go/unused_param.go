package demo

import "fmt"

func FindElement(x, y int) {
	nums := Int{1, 2, 3, 4, 5}
	if nums.Present(x) {
		fmt.Printf("%d is present in %v\n", x, nums)
	} else {
		fmt.Printf("%d is not present in %v\n", x, nums)
	}
}
