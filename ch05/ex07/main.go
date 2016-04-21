package main

import (
	"os"

	"./outline"
)

func main() {
	for _, url := range os.Args[1:] {
		outline.Outline(url)
	}
}
