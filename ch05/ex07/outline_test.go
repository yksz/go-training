package main

import (
	"bytes"
	"fmt"
	"testing"

	"golang.org/x/net/html"
)

func TestOutline(t *testing.T) {
	url := "http://gopl.io"

	var buf *bytes.Buffer
	buf = new(bytes.Buffer)
	out = buf
	outline(url)
	outlineHTML := buf.String()

	doc, err := html.Parse(buf)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	buf = new(bytes.Buffer)
	out = buf
	forEachNode(doc, startElement, endElement)
	parsedHTML := buf.String()

	if outlineHTML != parsedHTML {
		fmt.Println("!!! OUTLINE HTML !!!")
		fmt.Println(outlineHTML)
		fmt.Println("!!! PARSED HTML !!!")
		fmt.Println(parsedHTML)
		t.Errorf("outlineHTML != parsedHTML\n")
	}
}
