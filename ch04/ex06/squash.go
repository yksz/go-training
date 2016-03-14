package main

import "fmt"

func main() {
	s := "a\t\nb\v\fc\r d"
	fmt.Printf("%q\n", squashSpaces([]byte(s))) // "a b c d"
}

func squashSpaces(bytes []byte) []byte {
	out := bytes[:0]
	spaced := false
	for _, b := range bytes {
		if isSpace(b) {
			if !spaced {
				out = append(out, ' ')
				spaced = true
			}
		} else {
			out = append(out, b)
			spaced = false
		}
	}
	return out
}

func isSpace(b byte) bool {
	spaces := []byte{'\t', '\n', '\v', '\f', '\r', ' '}
	for _, space := range spaces {
		if b == space {
			return true
		}
	}
	return false
}
