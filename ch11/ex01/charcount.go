// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

var (
	in  io.Reader = os.Stdin
	out io.Writer = os.Stdout
)

func main() {
	charcount()
}

func charcount() {
	counts := make(map[string]int)  // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	bin := bufio.NewReader(in)
	for {
		r, n, err := bin.ReadRune() // returns rune, nbytes, error
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
		counts[string(r)]++
		utflen[n]++
	}
	fmt.Fprintf(out, "rune\tcount\n")
	for _, c := range sortedKeys(counts) {
		fmt.Fprintf(out, "'%s'\t%d\n", c, counts[c])
	}
	fmt.Fprint(out, "\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Fprintf(out, "%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Fprintf(out, "\n%d invalid UTF-8 characters\n", invalid)
	}
}

func sortedKeys(m map[string]int) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
