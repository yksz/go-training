package sexpr

import "fmt"

func Example() {
	type object struct {
		BoolT      bool
		BoolF      bool
		Float32    float32
		Float64    float64
		Complex64  complex64
		Complex128 complex128
		Interface  interface{}
	}
	obj := object{
		BoolT:      true,
		BoolF:      false,
		Float32:    1.2,
		Float64:    3.4,
		Complex64:  1 + 2i,
		Complex128: 3 + 4i,
		Interface:  []int{1, 2, 3},
	}

	data, err := MarshalIndent(obj)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", data)

	// Output:
	// ((BoolT t) (BoolF nil) (Float32 1.200000) (Float64 3.400000)
	//  (Complex64 #C(1.000000 2.000000)) (Complex128 #C(3.000000 4.000000))
	//  (Interface ("[]int" (1 2 3))))
}
