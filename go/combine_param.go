package demo

import "fmt"

func CombineParam(a int, b int, c int, d string, e string) {
	nums := Int{a, b, c}
	nums.Map(func(n int) int {
		return n * (len(d) + len(e))
	})
	fmt.Println(nums)
}
