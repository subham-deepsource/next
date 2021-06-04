package demo

import "time"

func Since() time.Duration {
	u := time.Now()
	time.Sleep(1 * time.Second)
	return time.Now().Sub(u)
}
