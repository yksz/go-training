package main

import "fmt"

func main() {
	s := "0 1 2 3 4 5"
	b := []byte(s)
	reverse(b)
	fmt.Printf("%q\n", b) // "5 4 3 2 1 0"
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
