package intset

import "fmt"

func ExampleElems() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	for _, v := range x.Elems() {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 9
	// 144
}
