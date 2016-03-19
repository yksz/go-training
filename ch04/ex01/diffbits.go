package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint8) int {
	return int(pc[byte(x)])
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	count := countDiffBits(c1, c2)
	fmt.Printf("%x\n%x\n", c1, c2)
	fmt.Printf("diff=%d\n", count)
}

func countDiffBits(c1 [32]byte, c2 [32]byte) int {
	count := 0
	for i := 0; i < sha256.Size; i++ {
		count += PopCount(c1[i] ^ c2[i])
	}
	return count
}
