package demo

func ExampleDeBlank() {
	DeBlank()
	// Output:
	// 3 from [1 2 3 4] is not present in [1 2 4 8]
}

func ExampleCombineAppend() {
	CombineAppend()
	// Output:
	// [1 2 3 4 5 6 7]
}

func ExampleCombineParam() {
	CombineParam(1, 2, 3, "x", "xy")
	// Output:
	// [3 6 9]
}

func ExampleRuneStr() {
	s := "deepsource"
	RuneStr(s)
	// Output:
	// deepsource
}

func ExampleFindElement() {
	FindElement(10, 0)
	FindElement(5, 0)
	// Output:
	// 10 is not present in [1 2 3 4 5]
	// 5 is present in [1 2 3 4 5]
}
