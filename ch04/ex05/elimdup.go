package main

import "fmt"

func main() {
	data := []string{"h", "e", "l", "l", "o"}
	fmt.Printf("%q\n", eliminateDuplicates(data)) // ["h" "e" "l" "o"]
}

func eliminateDuplicates(strings []string) []string {
	result := strings[:1]
	prev := strings[0]
	for _, s := range strings[1:] {
		if s != prev {
			result = append(result, s)
		}
		prev = s
	}
	return result
}
