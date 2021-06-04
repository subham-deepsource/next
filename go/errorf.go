package demo

import (
	"errors"
	"fmt"
)

func ErrorFormat() error {
	return errors.New(fmt.Sprintf("1+2 != %d", 4))
}
