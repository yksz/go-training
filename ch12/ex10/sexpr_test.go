package sexpr

import (
	"reflect"
	"testing"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v
//
func Test(t *testing.T) {
	type object struct {
		BoolT      bool
		BoolF      bool
		Float32    float32
		Float64    float64
		Interface1 interface{}
		Interface2 interface{}
	}
	obj := object{
		BoolT:      true,
		BoolF:      false,
		Float32:    1.2,
		Float64:    3.4,
		Interface1: 1,
		Interface2: []int{1, 2, 3},
	}

	// Encode it
	data, err := Marshal(obj)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var got object
	if err := Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal failed: %v:", err)
	}
	t.Logf("Unmarshal() = %+v\n", got)

	// Check equality.
	if !reflect.DeepEqual(got, obj) {
		t.Fatal("not equal")
	}
}
