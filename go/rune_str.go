package demo

import (
	"fmt"
)

func RuneStr(s string) {
	for _, r := range []rune(s) {
		fmt.Print(string(r))
	}
}
