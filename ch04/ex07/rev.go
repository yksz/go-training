package main

import "fmt"

func main() {
	s := "Hello, 世界"
	b := []byte(s)
	b = reverse(b)
	fmt.Printf("%q\n", b) // "界世 ,olleH"
}

func reverse(b []byte) []byte {
	r := []rune(string(b))
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return []byte(string(r))
}
