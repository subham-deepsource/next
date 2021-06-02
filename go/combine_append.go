package demo

import "fmt"

func CombineAppend() {
	slice := Int{1, 2, 3, 4}
	slice = append(slice, 5)
	slice = append(slice, 6)
	slice = append(slice, 7)

	fmt.Println(slice)
}
