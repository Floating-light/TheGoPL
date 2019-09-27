package main

import (
	"fmt"

	"github.com/Floating-light/TheGoPL/ch3"
)

func main() {
	fmt.Println(ch3.CommaNonRecursiveBuffer("83457029.345"))
	fmt.Println(ch3.CommaNonRecursiveBuffer("1"))
	fmt.Println(ch3.CommaNonRecursiveBuffer("10"))
	fmt.Println(ch3.CommaNonRecursiveBuffer("100"))
	fmt.Println(ch3.CommaNonRecursiveBuffer("1000"))
	fmt.Println(ch3.CommaNonRecursiveBuffer("10000"))
	fmt.Println(ch3.CommaNonRecursiveBuffer("100000"))
	fmt.Println(ch3.CommaNonRecursiveBuffer("1000000"))
	fmt.Println(ch3.CommaNonRecursiveBuffer("+100000.00"))
	fmt.Println(ch3.CommaNonRecursiveBuffer("-100000000"))
}
