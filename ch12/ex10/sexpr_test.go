package sexpr

import (
	"reflect"
	"testing"
)

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

	data, err := Marshal(obj)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var got object
	if err := Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal failed: %v:", err)
	}

	if !reflect.DeepEqual(got, obj) {
		t.Fatalf("got:%v\nwant:%v\n", got, obj)
	}
}
