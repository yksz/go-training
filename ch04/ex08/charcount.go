// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	categories := map[string]int{
		"letter": 0,
		"mark":   0,
		"number": 0,
		"punct":  0,
		"space":  0}
	counters := []func(rune){
		func(r rune) {
			if unicode.IsLetter(r) {
				categories["letter"]++
			}
		},
		func(r rune) {
			if unicode.IsMark(r) {
				categories["mark"]++
			}
		},
		func(r rune) {
			if unicode.IsNumber(r) {
				categories["number"]++
			}
		},
		func(r rune) {
			if unicode.IsPunct(r) {
				categories["punct"]++
			}
		},
		func(r rune) {
			if unicode.IsSpace(r) {
				categories["space"]++
			}
		}}

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		for _, counter := range counters {
			counter(r)
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\ncategory count\n")
	for k, v := range categories {
		fmt.Printf("%-8s %d\n", k, v)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
