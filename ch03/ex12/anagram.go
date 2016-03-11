package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	s := []string{"silent", "listen"}
	if len(os.Args) > 2 {
		s = os.Args[1:]
	}

	ans := "NOT anagram"
	if areAnagram(s) {
		ans = "anagram"
	}
	fmt.Printf("%s are %s\n", strings.Join(s, " and "), ans)
}

func areAnagram(s []string) bool {
	if len(s) <= 1 {
		return false
	}

	s = sortRuneSlice(s)
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			return false
		}
	}
	return true
}

func sortRuneSlice(s []string) []string {
	sorts := make([]string, 0, len(s))
	for _, v := range s {
		r := []rune(v)
		sort.Sort(RuneSlice(r))
		sorts = append(sorts, string(r))
	}
	return sorts
}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
