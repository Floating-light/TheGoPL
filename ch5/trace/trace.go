package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("big slow operation")()
	time.Sleep(10 * time.Second)
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

//------------------------------------------------------------------------------
//Deferred functions run after return statements have updated the fucntion's
//result variables.Because an anonymous function can access its enclosing
//function's variables, including named results, a deferred anonymous function
//can observe the function's results.
func double(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	}()
	return x + x
}

//A deferred anonymous function can even change the values that the enclosing
//fucntion returns to its caller.
func triple(x int) (result int) {
	defer func() {
		result += x
		fmt.Printf("triple(%d) = %d\n", x, result)
	}()
	return double(x)
}

func main() {
	fmt.Println(triple(10))
}
