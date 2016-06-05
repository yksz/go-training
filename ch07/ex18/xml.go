package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func Parse(r io.Reader) (Node, error) {
	dec := xml.NewDecoder(os.Stdin)
	var root, parent *Element
	var stack []*Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			elem := new(Element)
			elem.Type = tok.Name
			elem.Attr = tok.Attr
			if parent == nil {
				root = elem
			} else {
				parent.Children = append(parent.Children, elem)
			}
			parent = elem
			stack = append(stack, elem) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
			if len(stack) != 0 {
				parent = stack[len(stack)-1]
			}
		case xml.CharData:
			if parent != nil {
				parent.Children = append(parent.Children, CharData(tok))
			}
		}
	}
	return root, nil
}

func main() {
	doc, err := Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "xml: %v\n", err)
		os.Exit(1)
	}
	outline(doc)
}

func outline(doc Node) {
	var depth int
	forEachNode(doc,
		func(n Node) {
			switch n := n.(type) {
			case *Element:
				fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Type.Local, getAttrs(n))
				depth++
			case CharData:
				s := strings.TrimSpace(string(n))
				if s != "" {
					fmt.Printf("%*s%s\n", depth*2, "", s)
				}
			}
		},
		func(n Node) {
			switch n := n.(type) {
			case *Element:
				depth--
				fmt.Printf("%*s</%s>\n", depth*2, "", n.Type.Local)
			}
		})
}

func forEachNode(n Node, pre, post func(n Node)) {
	if pre != nil {
		pre(n)
	}

	if elem, ok := n.(*Element); ok {
		for _, c := range elem.Children {
			forEachNode(c, pre, post)
		}
	}

	if post != nil {
		post(n)
	}
}

func getAttrs(e *Element) string {
	s := ""
	for _, a := range e.Attr {
		if a.Value != "" {
			s += fmt.Sprintf(" %s=\"%s\"", a.Name.Local, a.Value)
		} else {
			s += fmt.Sprintf(" %s", a.Name.Local)
		}
	}
	return s
}
