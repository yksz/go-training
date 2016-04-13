package main

import (
	"fmt"
)

func join(sep string, s ...string) string {
	if len(s) == 0 {
		return ""
	}

	result := s[0]
	for _, v := range s[1:] {
		result += sep + v
	}
	return result
}

func main() {
	fmt.Println(join(", ", "foo", "bar", "baz")) // "foo, bar, baz"
}
