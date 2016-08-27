package cycle

import (
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2 CyclePtr
	cyclePtr1 = &cyclePtr1
	cyclePtr2 = &cyclePtr2

	type CycleSlice []CycleSlice
	var cycleSlice CycleSlice
	cycleSlice = append(cycleSlice, cycleSlice)

	for _, test := range []struct {
		x    interface{}
		want bool
	}{
		// basic types
		{1, false},
		{"foo", false},
		{nil, false},
		// slices
		{[]string{"foo"}, false},
		{[]string{}, false},
		// slice cycles
		{cycleSlice, true},
		// pointer cycles
		{cyclePtr1, true},
		{cyclePtr2, true},
	} {
		if IsCycle(test.x) != test.want {
			t.Errorf("IsCycle(%v) = %t", test.x, !test.want)
		}
	}
}

func Example_IsCycle() {
	// Circular linked lists a -> b -> a and c -> c.
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	fmt.Println(IsCycle(a)) // "true"
	fmt.Println(IsCycle(b)) // "true"
	fmt.Println(IsCycle(c)) // "true"

	// Output:
	// true
	// true
	// true
}
