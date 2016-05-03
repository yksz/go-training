package main

import (
	"bytes"
	"fmt"
	"os"

	"./count"
)

func main() {
	var b bytes.Buffer
	c, n := count.CountingWriter(&b)
	c.Write([]byte("Hello "))
	fmt.Fprintf(c, "world!\n")
	b.WriteTo(os.Stdout)
	fmt.Printf("%d\n", *n) // "13"
}
