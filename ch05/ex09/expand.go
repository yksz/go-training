package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "abc $foo def $bar"
	fmt.Println(expand(s, strings.ToUpper)) // "abc FOO def BAR"
}

func expand(s string, f func(string) string) string {
	var result, target []rune
	isTarget := false
	for _, r := range []rune(s) {
		if r == '$' {
			isTarget = true
			continue
		}
		if unicode.IsSpace(r) && isTarget {
			result = append(result, []rune(f(string(target)))...)
			target = target[:0] // clear
			isTarget = false
		}
		if isTarget {
			target = append(target, r)
		} else {
			result = append(result, r)
		}
	}
	result = append(result, []rune(f(string(target)))...)
	return string(result)
}
