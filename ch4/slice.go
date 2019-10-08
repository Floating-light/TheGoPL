//slice not comparable
package ch4

import "log"

//Reverse reverse string slice
func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]

	}
}

//Rotate a slice by n elements
//if direct is true, rotate to left, else rotate to right
func Rotate(s []int, n int, direct bool) {
	if n > len(s) {
		log.Fatalf("n > len(s)")
		return
	}
	if !direct {
		Reverse(s)
	}
	Reverse(s[:n])
	Reverse(s[n:])
	if direct {
		Reverse(s)
	}
}

//SliceCompare slice are not comparable, we must do it ourselves.
func SliceCompare(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

//Nonempty return a slice holding only the non-empty strings.
//The underlying array is modified during the call.
func Nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

//Nonempty2 using append
func Nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
