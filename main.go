package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Floating-light/TheGoPL/ch4"
)

func longestCommonPrefix(strs []string) string {
	num := len(strs)
	if num == 0 || len(strs[0]) == 0 {
		return ""
	}
	pre := ""
	for i := 0; i < len(strs[0]); i++ {
		cur := strs[0][0 : i+1]
		j := 1
		for ; j < num; j++ {
			if !strings.HasPrefix(strs[j], cur) {
				break
			}
		}
		if j != num {
			break
		} else {
			pre = cur
		}
	}
	return pre
}

func print() {
	fmt.Printf("Test timer print\n")
}
func startTimer(f func()) {
	go func() {
		for {
			f()
			now := time.Now()

			next := now.Add(time.Second * 5)
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

func main() {
	/* startTimer(print)
	fmt.Printf("end main function\n")
	for {

	} */
	ch4.FalseMain()
}
