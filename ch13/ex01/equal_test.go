package equal

import "testing"

func TestEqual(t *testing.T) {
	for _, test := range []struct {
		x, y interface{}
		want bool
	}{
		{1, 1, true},
		{1, 2, false},  // different values
		{1, 1.0, true}, // different types
		{1, 1.0 + 1.0e-10, false},
		{1, 1.0 + 0.9e-10, true},
	} {
		if Equal(test.x, test.y) != test.want {
			t.Errorf("Equal(%v, %v) = %t",
				test.x, test.y, !test.want)
		}
	}
}
