package sexpr

import "fmt"

func Example() {
	type object struct {
		Int    int
		Uint   uint
		String string
		Ptr    *int
		Array  [3]int
		Slice  []int
		Map    map[int]string
	}
	obj := object{Uint: 1}

	data, err := Marshal(obj)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", data)

	// Output:
	// ( (Uint 1) )
}
