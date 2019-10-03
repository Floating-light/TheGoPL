package ch4

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popCount(x [32]uint8) int {
	var count int = 0
	for i := range x {
		count += int(pc[i])
	}
	return count
}

//Sha256 Notice %x print all the elements of an array or slice
//of bytes in hexadecimal,%t to show a boolean, and %T to display the type of a value.
func Sha256() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

//Zero the contents of a [32]byte array
func Zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

//ZeroLiteral the contents of a [32]byte array using [32]byte array zero literals
func ZeroLiteral(ptr *[32]byte) {
	*ptr = [32]byte{}
}

//CountDiffBitsOfSha256 .
func CountDiffBitsOfSha256() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	fmt.Printf("c1 : %d, c2 : %d\n", popCount(c1), popCount(c2))
}
