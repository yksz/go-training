package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Printf("usage: %s <element id> <urls>\n", os.Args[0])
		os.Exit(1)
	}
	id := os.Args[1]
	for _, url := range os.Args[2:] {
		elem, err := fetchAndGetElement(url, id)
		if err != nil {
			log.Fatal(err)
		}
		printElement(elem)
	}
}

func printElement(elem *html.Node) {
	fmt.Printf("<%s", elem.Data)
	for _, a := range elem.Attr {
		fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
	}
	fmt.Println(">")
}

func fetchAndGetElement(url, id string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return ElementByID(doc, id), nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var elem *html.Node
	forEachNode(doc,
		func(n *html.Node) bool {
			if n.Type == html.ElementNode {
				for _, a := range n.Attr {
					if a.Key == "id" && a.Val == id {
						elem = n
						return false
					}
				}
			}
			return true
		},
		nil)
	return elem
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil {
		if !pre(n) {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		if !post(n) {
			return
		}
	}
}
