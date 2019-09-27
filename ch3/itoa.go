package ch3

import (
	"fmt"
	"strconv"
)

func ConvStrInteger() {
	x := 12345
	y := fmt.Sprintf("%d", x)       //convert by fmt.Sprintf
	fmt.Println(y, strconv.Itoa(x)) //convert by strconv.Itoa

	//format numbers in different base by FormatInt and FormatUint
	fmt.Println(strconv.FormatInt(int64(x), 2))
	fmt.Println(strconv.FormatInt(int64(x), 8))
	fmt.Println(strconv.FormatInt(int64(x), 10))
	fmt.Println(strconv.FormatInt(int64(x), 16))

	//fromat by fmt.Sprintf, using %b, %d, %u, %x
	s1 := fmt.Sprintf("x = %b", x)
	s2 := fmt.Sprintf("x = %d", x)
	s3 := fmt.Sprintf("x = %u", x)
	s4 := fmt.Sprintf("x = %x", x)

	fmt.Println(s1, s2, s3, s4)

}
