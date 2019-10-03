package ch3

import (
	"fmt"
	"time"
)

func Constants() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

	const (
		a = 1
		b
		c = 100
		d
	)
	fmt.Println(a, b, c, d)
}
