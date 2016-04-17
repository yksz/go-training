package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "\t\n\va\f\rb c　"
	fmt.Printf("%q\n", squashSpaces([]byte(s))) // " a b c "
}

func squashSpaces(b []byte) []byte {
	runes := []rune(string(b))
	result := runes[:0]
	spaced := false
	for _, r := range runes {
		if unicode.IsSpace(r) {
			if !spaced {
				result = append(result, ' ')
				spaced = true
			}
		} else {
			result = append(result, r)
			spaced = false
		}
	}
	return []byte(string(result))
}
