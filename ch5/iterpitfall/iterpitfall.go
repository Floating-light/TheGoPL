package main

import "fmt"

func caveat() {
	var delayExcute []func()
	contents := []string{"test1", "test2", "test3", "test4", "test5", "test6"}
	for _, content := range contents {
		content := content //surprising
		delayExcute = append(delayExcute, func() {
			fmt.Printf("%s\n", content)
		})
	}
	for _, exc := range delayExcute {
		exc()
	}
}

func main() {
	caveat()
}
