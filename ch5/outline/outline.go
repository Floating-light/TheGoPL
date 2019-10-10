package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//outline 尽管outline会向stack中添加元素，改变底层数组，但是对调用者outline而言，
//这并没有改变对调用者可见的初始元素，所以当被调函数返回时，调用者的stack还和调用之前的一样。
//关键在于slice被copy时，仅仅底层数组是指针copy，会影响外层的slice，其它的size，
//capabilities并没有共享，调用函数内外是不同的。
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}
