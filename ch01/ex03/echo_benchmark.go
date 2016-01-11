package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%f s\n", time.Since(start).Seconds())
}

func echo2() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%f s\n", time.Since(start).Seconds())
}

func main() {
	echo1()
	echo2()
}
