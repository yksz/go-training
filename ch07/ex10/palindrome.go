package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < s.Len() || j > 0; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) { // equal
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	{
		values := []int{3, 1, 4, 1}
		fmt.Println(IsPalindrome(sort.IntSlice(values))) // "false"
	}
	{
		values := []int{3, 1, 4, 1, 3}
		fmt.Println(IsPalindrome(sort.IntSlice(values))) // "true"
	}
}
