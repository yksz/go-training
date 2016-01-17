package popcount

import (
	"reflect"
	"testing"
)

func assert(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("(expected, actual) = (%v, %v)\n", expected, actual)
	}
}

func TestPopCount(t *testing.T) {
	assert(t, 32, PopCount(0x1234567890ABCDEF))
}

func TestPopCountByBitClear(t *testing.T) {
	assert(t, 32, PopCountByBitClear(0x1234567890ABCDEF))
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByBitClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByBitClear(0x1234567890ABCDEF)
	}
}
