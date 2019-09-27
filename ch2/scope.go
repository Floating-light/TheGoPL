package ch2

import "fmt"

func f() {
	fmt.Println("This funciton f()")
}

//var f int <f redeclared in this block
//previous declaration at .\scope.go:5:6g>
var true string = "true"

func hideOuter() {
	f := "hide outer function f() "
	fmt.Println(f)
	//f() Error: cannot call non-function f (type string)
}
