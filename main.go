package main

import (
	"fmt"

	"github.com/Floating-light/TheGoPL/ch2"
)

func main() {
	fmt.Println(ch2.BoilingC - ch2.FreezingC)
	fmt.Println(ch2.CToF(ch2.BoilingC - ch2.FreezingC))
	fmt.Println(ch2.CToF(ch2.Celsius(float64(float32(int(37))))))
}
