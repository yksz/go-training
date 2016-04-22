// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var out io.Writer = os.Stdout

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild != nil {
			fmt.Fprintf(out, "%*s<%s%s>\n", depth*2, "", n.Data, getAttrs(n))
			if n.FirstChild.Type == html.TextNode {
				s := strings.TrimSpace(n.FirstChild.Data)
				if s != "" {
					fmt.Fprintf(out, "%s\n", s)
				}
			}
		} else {
			fmt.Fprintf(out, "%*s<%s%s>\n", depth*2, "", n.Data, getAttrs(n))
		}
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func getAttrs(n *html.Node) string {
	s := ""
	for _, a := range n.Attr {
		if a.Val != "" {
			s += fmt.Sprintf(" %s=\"%s\"", a.Key, a.Val)
		} else {
			s += fmt.Sprintf(" %s", a.Key)
		}
	}
	return s
}
