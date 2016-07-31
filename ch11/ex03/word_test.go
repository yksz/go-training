package word

import (
	"math/rand"
	"testing"
	"time"
)

func randomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(23) + 2 // random length: 2-24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		randN := rng.Intn(93) + 33 // ascii: 33-125
		runes[i] = rune(randN)
		runes[n-1-i] = rune(randN + 1)
	}
	return string(runes)
}

func TestRandomNonPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}
