package demo

import (
	"os"
)

func HostnameErrNil() (string, error) {
	name, err := os.Hostname()
	if err == nil {
		return name, err
	}
	return "", err
}
