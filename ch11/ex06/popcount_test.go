package popcount_test

import (
	"testing"

	pc4 "../../ch02/ex04"
	pc5 "../../ch02/ex05"
	pc "gopl.io/ch2/popcount"
)

func BenchmarkPopCount10000(b *testing.B) {
	benchmarkPopCount(b, 10000)
}

func BenchmarkPopCount100000(b *testing.B) {
	benchmarkPopCount(b, 100000)
}

func BenchmarkPopCount1000000(b *testing.B) {
	benchmarkPopCount(b, 1000000)
}

func benchmarkPopCount(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		pc.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByBitShift10000(b *testing.B) {
	benchmarkPopCountByBitShift(b, 10000)
}

func BenchmarkPopCountByBitShift100000(b *testing.B) {
	benchmarkPopCountByBitShift(b, 100000)
}

func BenchmarkPopCountByBitShift1000000(b *testing.B) {
	benchmarkPopCountByBitShift(b, 1000000)
}

func benchmarkPopCountByBitShift(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		pc4.PopCountByBitShift(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByBitClear10000(b *testing.B) {
	benchmarkPopCountByBitShift(b, 10000)
}

func BenchmarkPopCountByBitClear100000(b *testing.B) {
	benchmarkPopCountByBitShift(b, 100000)
}

func BenchmarkPopCountByBitClear1000000(b *testing.B) {
	benchmarkPopCountByBitShift(b, 1000000)
}

func benchmarkPopCountByBitClear(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		pc5.PopCountByBitClear(0x1234567890ABCDEF)
	}
}
