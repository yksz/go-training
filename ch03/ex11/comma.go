// Comma prints its argument numbers with a comma at each power of 1000.
package main

import (
	"fmt"
	"os"
	"strings"
)

const digit = 3

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", commaFloat(os.Args[i]))
	}
}

func commaFloat(s string) string {
	if len(s) <= digit {
		return s
	}

	sign := ""
	if first := s[:1]; first == "+" || first == "-" {
		sign = first
		s = s[1:]
	}
	if index := strings.Index(s, "."); index != -1 {
		return sign + comma(s[:index]) + s[index:]
	}
	return sign + comma(s)
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= digit {
		return s
	}
	return comma(s[:n-digit]) + "," + s[n-digit:]
}
