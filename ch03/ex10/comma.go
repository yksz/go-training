// Comma prints its argument numbers with a comma at each power of 1000.
package main

import (
	"bytes"
	"fmt"
	"os"
)

const digit = 3

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= digit {
		return s
	}

	var buf bytes.Buffer
	if mod := n % digit; mod != 0 {
		buf.WriteString(s[:mod])
		buf.WriteString(",")
		s = s[mod:]
	}
	buf.WriteString(s[:digit])
	for i := 1; i < n/digit; i++ {
		buf.WriteString(",")
		buf.WriteString(s[i*digit : (i+1)*digit])
	}
	return buf.String()
}
