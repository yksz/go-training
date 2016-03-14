// Rot rotates a slice.
package main

import "fmt"

func main() {
	{
		s := []int{0, 1, 2, 3, 4, 5}
		rotate(s, 2)
		fmt.Println(s) // "[2 3 4 5 0 1]"
	}
	{
		s := []int{0, 1, 2, 3, 4, 5}
		rotate(s, 8)
		fmt.Println(s) // "[2 3 4 5 0 1]"
	}
	{
		s := []int{0, 1, 2, 3, 4, 5}
		rotate(s, -2)
		fmt.Println(s) // "[4 5 0 1 2 3]"
	}
}

func rotate(s []int, left int) {
	count := left % len(s)
	if count < 0 {
		count += len(s)
	}

	for i := 0; i < count; i++ {
		for j := 0; j < len(s)-1; j++ {
			s[j], s[j+1] = s[j+1], s[j]
		}
	}
}
