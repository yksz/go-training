package sexpr

import "fmt"

func Example1() {
	type object struct {
		Bool      bool
		Float32   float32
		Complex64 complex64
		Interface interface{}
	}
	obj := object{
		Bool:      true,
		Float32:   1.2,
		Complex64: 1 + 2i,
		Interface: []int{1, 2, 3},
	}

	data, err := Marshal(obj)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", data)

	// Output:
	// ((Bool t) (Float32 1.200000) (Complex64 #C(1.000000 2.000000)) (Interface ("[]int" (1 2 3))))
}

func Example2() {
	type object struct {
		Bool       bool
		Float64    float64
		Complex128 complex128
		Interface  interface{}
	}
	obj := object{
		Bool:       false,
		Float64:    3.4,
		Complex128: 3 + 4i,
		Interface:  []string{"1", "2", "3"},
	}

	data, err := Marshal(obj)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", data)

	// Output:
	// ((Bool nil) (Float64 3.400000) (Complex128 #C(3.000000 4.000000)) (Interface ("[]string" ("1" "2" "3"))))
}
