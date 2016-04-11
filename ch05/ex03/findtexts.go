package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findtexts: %v\n", err)
		os.Exit(1)
	}
	visit(doc, nil)
}

func visit(n *html.Node, parent *html.Node) {
	if parent != nil && parent.Type == html.ElementNode {
		if parent.Data == "script" || parent.Data == "style" {
			return
		}
	}
	if n.Type == html.TextNode {
		if strings.TrimSpace(n.Data) != "" {
			fmt.Println(n.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, n)
	}
}
