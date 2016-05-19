// Package eval provides an expression evaluator.
package eval

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	if _, ok := env[v]; !ok {
		f := promptVar(v)
		env[v] = f
	}
	return env[v]
}

func promptVar(v Var) float64 {
	var in = bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s=", v)
		line, _, err := in.ReadLine()
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		f, err := strconv.ParseFloat(string(line), 64)
		if err != nil {
			continue
		}
		return f
	}
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
