package ch3

import (
	"fmt"
)

//NumerOverflow if the result of an arithmetic operation has more bits than can
//be represented int the rsult type, it is said to overflow. The high-order bits
// that do not fit are silently discarded.
func NumerOverflow() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) //u + 1 = 0001 0000 0000,u*u = 1111 1110 0000 0001

	var i int8 = 127
	//i=0111 1111, i+1 =1000 0000 = -128, 1111 1111 = -127
	fmt.Printf("%08b,\n%08b,\n%08b\n", i, i+1, i*i)
}

//BitWiseOper show how bitwise operation can be sued to interpret a uint8 value
// as a compact and efficient set of 8 independent bits.
func BitWiseOper() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) //0010 0010, set{1,5}
	fmt.Printf("%08b\n", y) //0000 0110, set{1,2}

	fmt.Printf("%08b\n", x&y)  //0000 0010
	fmt.Printf("%08b\n", x|y)  //0010 0110
	fmt.Printf("%08b\n", x^y)  //0010 0100
	fmt.Printf("%08b\n", x&^y) //0010 0000

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1) //0100 0100
	fmt.Printf("%08b\n", x>>1) //0001 0001
}

//SignedShift sign int shift test.
func SignedShift() {
	var x int8 = 1<<1 | 1<<5 // 0010 0010
	x = -x
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", x>>1)
	fmt.Printf("%08b\n", x>>2)

}

//ExpConv binary operators for arithmetic and logic (except shifts) must have
//operands of the same type.
func ExpConv() {
	var apples int32 = 1
	var oranges int16 = 2
	var _ int = int(apples + int32(oranges)) //invalid operation: apples + oranges (mismatched types int32 and int16)
}

//PrintNum print octal number, decimal numbers, hexadeciaml numbers
func PrintNum() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
}

//RunesChar rune, a symbol that has a mysterious or magic meaning.
func RunesChar() {
	ascii := 'a'
	unicode := '饕'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]c %[1]q\n", newline)
}
