// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func NewReader(s string) io.Reader {
	return &stringReader{s: s}
}

type stringReader struct {
	s     string
	index int64
}

func (r *stringReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if r.index >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n := copy(p, r.s[r.index:])
	r.index += int64(n)
	return n, nil
}

func main() {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc, err := html.Parse(NewReader(s))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	// Output:
	// foo
	// /bar/baz
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
