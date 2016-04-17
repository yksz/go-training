package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elems: %v\n", err)
		os.Exit(1)
	}
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	for _, h := range headings {
		fmt.Printf("<%s>", h.Data)
		if h.FirstChild != nil {
			fmt.Printf("%s", h.FirstChild.Data)
		}
		fmt.Printf("</%s>\n", h.Data)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	if len(name) == 0 {
		return nil
	}
	return visit(nil, doc, name)
}

func visit(elems []*html.Node, n *html.Node, tags []string) []*html.Node {
	for _, tag := range tags {
		if n.Type == html.ElementNode && n.Data == tag {
			elems = append(elems, n)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elems = visit(elems, c, tags)
	}
	return elems
}
