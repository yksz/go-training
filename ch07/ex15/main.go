package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"./eval"
)

var env = eval.Env{"PI": math.Pi}

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _, err := in.ReadLine()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		if len(line) == 0 {
			continue
		}
		expr, err := eval.Parse(string(line))
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("%.6g\n", expr.Eval(env))
	}
}
