package ch4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

//map define: map[K]V, there K is the type of key, V is the type of value.
//K must can be compare with "=="

func begin() {
	_ = make(map[string]int) //define, eq map[string]int{}
	b := map[string]int{
		"key": 23,
	}

	//remove
	delete(b, "key")
}

//CharCount count the rune from std io
func CharCount() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount : %v \n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for r, c := range counts {
		fmt.Printf("%q\t%d\n", r, c)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//Graph directed graph
var graph map[string]map[string]bool

//AddEdge add a edge
func AddEdge(from, to string) {
	edge := graph[from]
	if edge == nil {
		edge = make(map[string]bool)
		graph[from] = edge
	}
	edge[to] = true
}

//HasEdge .
func HasEdge(from, to string) bool {
	return graph[from][to]
}
