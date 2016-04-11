package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findelems: %v\n", err)
		os.Exit(1)
	}
	for elem, count := range visit(make(map[string]int), doc) {
		fmt.Printf("%s\t%d\n", elem, count)
	}
}

func visit(nelems map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		nelems[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(nelems, c)
	}
	return nelems
}
