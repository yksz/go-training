// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

type Attr map[string]string

type Element struct {
	Name string
	Attr Attr
}

func main() {
	selects := parseArgs()

	dec := xml.NewDecoder(os.Stdin)
	var stack []*Element // stack of elements
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			elem := &Element{Name: tok.Name.Local, Attr: make(Attr)}
			for _, attr := range tok.Attr {
				elem.Attr[attr.Name.Local] = attr.Value
			}
			stack = append(stack, elem) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, selects) {
				s := ""
				for _, elem := range stack {
					s += elem.Name + " "
					for k, v := range elem.Attr {
						s += k + "=" + v + " "
					}
				}
				fmt.Printf("%s: %s\n", s, tok)
			}
		}
	}
}

func parseArgs() []*Element {
	var selects []*Element
	for _, arg := range os.Args[1:] {
		if strings.Contains(arg, "=") { // attribute
			if len(selects) == 0 {
				continue
			}
			elem := selects[len(selects)-1]
			s := strings.Split(arg, "=")
			elem.Attr[s[0]] = s[1]
		} else {
			elem := &Element{Name: arg, Attr: make(Attr)}
			selects = append(selects, elem)
		}
	}
	return selects
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []*Element) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0].Name == y[0].Name &&
			(len(y[0].Attr) == 0 || reflect.DeepEqual(x[0].Attr, y[0].Attr)) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
