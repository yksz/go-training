package intset_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"gopl.io/ch6/intset"
)

var randN int

func init() {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	largeN := int32(math.Pow(2, 30))
	randN = int(rng.Int31n(largeN) + largeN)
	fmt.Printf("Random number: %d\n", randN)
}

func BenchmarkHas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x intset.IntSet
		x.Add(randN)
		x.Has(randN)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x intset.IntSet
		x.Add(randN)
	}
}

func BenchmarkUnionWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x, y intset.IntSet
		x.Add(randN)
		y.Add(randN)
		x.UnionWith(&y)
	}
}
