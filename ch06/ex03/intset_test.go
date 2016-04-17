package intset

import "fmt"

func ExampleIntersectWith() {
	{
		var x, y IntSet
		x.Add(1)
		x.Add(144)
		x.Add(9)
		y.Add(9)
		y.Add(42)
		x.IntersectWith(&y)
		fmt.Println(x.String()) // "{9}"
	}

	{
		var x, y IntSet
		x.Add(1)
		x.Add(144)
		x.Add(9)
		y.Add(9)
		y.Add(42)
		y.IntersectWith(&x)
		fmt.Println(y.String()) // "{9}"
	}

	// Output:
	// {9}
	// {9}
}

func ExampleDifferenceWith() {
	{
		var x, y IntSet
		x.Add(1)
		x.Add(144)
		x.Add(9)
		y.Add(9)
		y.Add(42)
		x.DifferenceWith(&y)
		fmt.Println(x.String()) // "{1 144}"
	}

	{
		var x, y IntSet
		x.Add(1)
		x.Add(144)
		x.Add(9)
		y.Add(9)
		y.Add(42)
		y.DifferenceWith(&x)
		fmt.Println(y.String()) // "{42}"
	}

	// Output:
	// {1 144}
	// {42}
}

func ExampleSymmetricDifference() {
	{
		var x, y IntSet
		x.Add(1)
		x.Add(144)
		x.Add(9)
		y.Add(9)
		y.Add(42)
		x.SymmetricDifference(&y)
		fmt.Println(x.String()) // "{1 42 144}"
	}

	{
		var x, y IntSet
		x.Add(1)
		x.Add(144)
		x.Add(9)
		y.Add(9)
		y.Add(42)
		y.SymmetricDifference(&x)
		fmt.Println(y.String()) // "{1 42 144}"
	}

	// Output:
	// {1 42 144}
	// {1 42 144}
}
