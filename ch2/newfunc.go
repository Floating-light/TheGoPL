package ch2

import (
	"fmt"
)

//NewFunc new function new(T)
//--> create a unnamed variable of type T
//--> initializes it to the zero value of T
//--> returns its address,which is a value of type *T
//new is no different from an ordinary local variable whose address is taken
func NewFunc() {
	p := new(int)
	fmt.Println(*p)
	fmt.Println(p)
	*p = 1008611
	fmt.Println(*p)
}
