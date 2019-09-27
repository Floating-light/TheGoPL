package ch3

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

//GoString .rune P70
func GoString() {
	s := "hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}

//BaseName1 remove directory components and a .suffix
//e.g., a=>a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func BaseName1(s string) string {
	//Discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	//Preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

//BaseName2 same as above, instead of using strings.LastIndex library function
func BaseName2(s string) string {
	slash := strings.LastIndex(s, "/") //-1, if slash not found
	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot != -1 {
		s = s[:dot]
	}
	return s
}

// Comma inserts commas in a non-negative decimal integer string.
func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

//CommaNonRecursiveBuffer .
func CommaNonRecursiveBuffer(s string) string {
	var sign string = ""
	var point string = ""
	if s[0] == '-' || s[0] == '+' {
		sign = s[:1]
		s = s[1:]
	}
	if i := strings.Index(s, "."); i != -1 {
		point = s[i:]
		s = s[:i]
	}
	var buf bytes.Buffer
	if r := len(s) % 3; r != 0 {
		buf.WriteString(s[:r])
		s = s[r:]
	}
	for len(s) >= 3 {
		if buf.Len() > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s[:3])
		s = s[3:]
	}
	return sign + buf.String() + point
}
