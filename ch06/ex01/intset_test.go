package intset

import "fmt"

func ExampleLen() {
	var x IntSet
	x.Add(1)
	fmt.Println(x.Len()) // "1"

	x.Add(144)
	fmt.Println(x.Len()) // "2"

	x.Add(9)
	fmt.Println(x.Len()) // "3"

	x.Add(1)
	fmt.Println(x.Len()) // "3"

	// Output:
	// 1
	// 2
	// 3
	// 3
}

func ExampleRemove() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	x.Remove(9)
	fmt.Println(x.String()) // "{1 144}"

	x.Remove(9)
	fmt.Println(x.String()) // "{1 144}"

	// Output:
	// {1 9 144}
	// {1 144}
	// {1 144}
}

func ExampleClear() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	x.Clear()
	fmt.Println(x.String()) // "{}"

	// Output:
	// {1 9 144}
	// {}
}

func ExampleCopy() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	y := x.Copy()
	y.Add(9)
	y.Add(42)
	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println(y.String()) // "{1 9 42 144}"

	// Output:
	// {1 9 144}
	// {1 9 42 144}
}
