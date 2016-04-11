// Findlinks prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var queries = map[string]string{
	"a":      "href",
	"img":    "src",
	"script": "src",
	"link":   "href",
	"iframe": "src",
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	for elemquery, attrquery := range queries {
		if n.Type == html.ElementNode && n.Data == elemquery {
			for _, a := range n.Attr {
				if a.Key == attrquery {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
