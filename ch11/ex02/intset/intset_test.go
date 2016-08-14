package intset

import (
	"testing"
)

func TestAddAndHas(t *testing.T) {
	var tests = []struct {
		input int
		want  bool
	}{
		{1, true},
		{144, true},
		{9, true},
		{123, false},
	}

	sets := []Set{&IntSet{}, NewMapIntSet()}
	for _, set := range sets {
		set.Add(1)
		set.Add(144)
		set.Add(9)
	}

	for _, set := range sets {
		for _, test := range tests {
			if got := set.Has(test.input); got != test.want {
				t.Errorf("set.Has(%d) = %v", test.input, got)
			}
		}
	}
}
