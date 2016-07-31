package intset_test

import (
	"testing"

	"gopl.io/ch6/intset"
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

	var x intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	for _, test := range tests {
		if got := x.Has(test.input); got != test.want {
			t.Errorf("x.Has(%d) = %v", test.input, got)
		}
	}
}
