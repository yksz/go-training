package split_test

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		s    string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		{"a:b:c", " ", 1},
		{"a:b:c", "", 5},
		{"a:b:", ":", 3},
		{":", ":", 2},
		{"", ":", 1},
	}

	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, got, test.want)
		}
	}
}
