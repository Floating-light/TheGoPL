package main

import "fmt"

//sum the type of vals is []int
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(0, 1, 2, 2, 3, 3, 45))
	fmt.Println(sum(1))
	fmt.Println(sum(10, 2, 6, 7))
	v := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(v...))

	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)
}

//but type of variadic function is distinct from the type of a functions
//with an ordinary slice parameter
func f(...int) {}
func g([]int)  {}
