package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	size := flag.Int("size", 256, "384 or 512 bits")
	flag.Parse()

	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "sha2: %v\n", err)
		os.Exit(1)
	}
	line = line[:len(line)-1]
	hashBySHA(line, *size)
}

func hashBySHA(s string, size int) {
	switch size {
	case 384:
		fmt.Printf("%x\n", sha512.Sum384([]byte(s)))
	case 512:
		fmt.Printf("%x\n", sha512.Sum512([]byte(s)))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(s)))
	}
}
