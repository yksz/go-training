package intset_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"../ex02/intset"
)

var randN int

func init() {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	largeN := int32(math.Pow(2, 30))
	randN = int(rng.Int31n(largeN) + largeN)
	fmt.Printf("Random number: %d\n", randN)
}

func BenchmarkIntSet_Has(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x intset.IntSet
		x.Has(randN)
	}
}

func BenchmarkIntSet_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x intset.IntSet
		x.Add(randN)
	}
}

func BenchmarkIntSet_UnionWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x, y intset.IntSet
		x.UnionWith(&y)
	}
}

func BenchmarkMapIntSet_Has(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := intset.NewMapIntSet()
		x.Has(randN)
	}
}

func BenchmarkMapIntSet_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := intset.NewMapIntSet()
		x.Add(randN)
	}
}

func BenchmarkMapIntSet_UnionWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := intset.NewMapIntSet()
		y := intset.NewMapIntSet()
		x.UnionWith(y)
	}
}
