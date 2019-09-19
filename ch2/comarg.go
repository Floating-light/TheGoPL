package ch2

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")
var ser = flag.String("server", " balabala", "server path")
var mode string

//CommandLineFlag echo4 print command line arguments
func CommandLineFlag() {
	flag.StringVar(&mode, "m", "debug", "set running model, debug or relese")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	fmt.Printf("server ip arguments : %s\n", *ser)
	fmt.Printf("mode : %s\n", mode)
	if !*n {
		fmt.Println()
	}
}
