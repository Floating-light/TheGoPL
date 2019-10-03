package main

import (
	"fmt"

	"github.com/Floating-light/TheGoPL/ch4"
)

func main() {
	a := [...]int{1, 2, 3, 4}
	fmt.Println(a)
	ch4.Rotate(a[:], 2, true)
	fmt.Println(a)

	b := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(b)
	ch4.Rotate(b[:], 4, false)
	fmt.Println(b)

	sli := []string{"test", "", "adghtr"}

	fmt.Println(ch4.Nonempty(sli))
	fmt.Println(sli)
}
