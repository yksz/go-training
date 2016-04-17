package intset

import "fmt"

func ExampleAddAll() {
	var x IntSet
	x.AddAll()
	fmt.Println(x.String()) // "{}"

	x.AddAll(1, 2, 3)
	fmt.Println(x.String()) // "{1 2 3}"

	x.AddAll(4, 5, 6)
	fmt.Println(x.String()) // "{1 2 3 4 5 6}"

	// Output:
	// {}
	// {1 2 3}
	// {1 2 3 4 5 6}
}
