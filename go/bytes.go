package demo

import "bytes"

func bytesEq() bool {
	a, b := []byte("a"), []byte("b")
	return bytes.Compare(a, b) == 0
}
