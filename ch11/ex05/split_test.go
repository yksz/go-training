package split_test

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		s    string
		sep  string
		want []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"a:b:c", " ", []string{"a:b:c"}},
		{"a:b:c", "", []string{"a", ":", "b", ":", "c"}},
		{"a:b:", ":", []string{"a", "b", ""}},
		{":", ":", []string{"", ""}},
		{"", ":", []string{""}},
	}

	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got := words; !reflect.DeepEqual(got, test.want) {
			t.Errorf("Split(%q, %q) returned %q, want %q", test.s, test.sep, got, test.want)
		}
	}
}
