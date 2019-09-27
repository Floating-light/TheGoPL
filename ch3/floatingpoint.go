package ch3

import (
	"fmt"
	"math"
)

//AccumError accumulate error of float32 ,be careful, use float64 instead.
func AccumError() {
	var f float32 = 16777216 // 1<< 24
	fmt.Println(f == f+1)
	var f64 float64 = 1 << 50
	fmt.Println(f64 == f64+1)
}

//PrintControll the print foramt controll
func PrintControll() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e……x = %8.3f\n", x, math.Exp(float64(x)))
	}
}

//SpecialValue positive and negative infinities, excessive magnitude(极值),
//division by zero, NaN("Not a Number"), 0/0, Sqrt(-1)
func SpecialValue() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
	fmt.Println(math.NaN(), math.IsNaN(z/z), math.Inf(0))

	//anything test with NaN is false
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan, nan > 0, nan == 3.14, (nan+10)/2)
}

//for return a floating-point result might fail, it's better to report
//the failure separately, like this:
/*
func compute() (value float64, ok bool) {
	// ...
	if failed {
		return 0, false
	}
	return result, true
}
*/
