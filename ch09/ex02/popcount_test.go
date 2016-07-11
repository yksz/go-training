package popcount_test

import (
	"testing"

	"./popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		count := popcount.PopCount(0x1234567890ABCDEF)
		if count != 32 {
			b.Errorf("count: %f\n", count)
		}
	}
}
