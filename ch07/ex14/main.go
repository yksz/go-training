package main

import (
	"fmt"
	"math"
	"os"

	"./eval"
)

var env = eval.Env{"PI": math.Pi}

func main() {
	// macro computes the minimum value of its operands
	s := "macro(1, pow(0.5, 2), sin(PI/2), sqrt(2))"
	expr, err := eval.Parse(s)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%.6g\n", expr.Eval(env)) // "0.25"
}
