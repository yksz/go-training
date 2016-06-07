package main

import (
	"fmt"
	"os"

	"./eval"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("usage: %s <expression>\n", os.Args[0])
		os.Exit(1)
	}

	expr1, err := eval.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("expr1: %s\n", expr1)

	expr2, err := eval.Parse(expr1.String())
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("expr2: %s\n", expr2)

	if expr1.String() != expr2.String() {
		fmt.Println("expr1 != expr2")
	}
	fmt.Println("expr1 == expr2")
}
