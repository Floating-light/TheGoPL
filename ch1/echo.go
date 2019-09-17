package ch1

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//Exe1 print commmand-line arguments ,using strings.Jion
func Exe1() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

//Inefficient print command-line arguments, by for loop
func Inefficient() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

//Exe2 print command-line arguments, range for loop
func Exe2() {
	for i, arg := range os.Args {
		fmt.Printf("The number %d argument is : %s\n", i, arg)
	}
}

//Exe3 measure the difference in runing time between inefficient version and the one that uses string.Jion.
func Exe3() {
	startFast := time.Now()
	Exe1()
	fast := time.Since(startFast).Nanoseconds()

	startSlowly := time.Now()
	Inefficient()
	slowly := time.Since(startSlowly).Nanoseconds()

	fmt.Printf("Fast version time cost : %dns\nSlowly version time cost : %dns\n", fast, slowly)
}
