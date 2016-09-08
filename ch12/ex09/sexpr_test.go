package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Example() {
	data := `((Title "Dr. Strangelove")` +
		` (Subtitle "How I Learned to Stop Worrying and Love the Bomb")` +
		` (Year 1964))`

	dec := NewDecoder(bytes.NewReader([]byte(data)))
	var slice []string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Decode failed: %v", err)
		}
		var v string
		switch tok := tok.(type) {
		case Symbol:
			v = string(tok)
			fmt.Printf("symbol: %s\n", v)
		case String:
			v = string(tok)
			fmt.Printf("string: %s\n", v)
		case Int:
			v = strconv.Itoa(int(tok))
			fmt.Printf("int: %s\n", v)
		case StartList:
			v = string(tok)
			fmt.Printf("start: %s\n", v)
		case EndList:
			v = string(tok)
			fmt.Printf("end: %s\n", v)
		}
		slice = append(slice, v)
	}
	fmt.Printf("%s\n", strings.Join(slice, " "))

	// Output:
	// start: (
	// start: (
	// symbol: Title
	// string: "Dr. Strangelove"
	// end: )
	// start: (
	// symbol: Subtitle
	// string: "How I Learned to Stop Worrying and Love the Bomb"
	// end: )
	// start: (
	// symbol: Year
	// int: 1964
	// end: )
	// end: )
	// ( ( Title "Dr. Strangelove" ) ( Subtitle "How I Learned to Stop Worrying and Love the Bomb" ) ( Year 1964 ) )
}
