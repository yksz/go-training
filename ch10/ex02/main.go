package main

import (
	"fmt"
	"os"

	"./archive"
	_ "./archive/tar"
	_ "./archive/zip"
)

func main() {
	for _, arg := range os.Args[1:] {
		a, kind, err := archive.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not open %s: %v\n", arg, err)
			continue
		}
		fmt.Printf("%s: %s\n", kind, arg)
		for _, file := range a.Files() {
			fmt.Printf(" - %s\n", file)
		}
	}
}
