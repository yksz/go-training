package main

import (
	"fmt"
	"math"
)

func max(vals ...float64) float64 {
	if len(vals) == 0 {
		return 0
	}

	max := math.SmallestNonzeroFloat64
	for _, val := range vals {
		max = math.Max(val, max)
	}
	return max
}

func min(vals ...float64) float64 {
	if len(vals) == 0 {
		return 0
	}

	min := math.MaxFloat64
	for _, val := range vals {
		min = math.Min(val, min)
	}
	return min
}

func max2(vals ...float64) float64 {
	if len(vals) < 1 {
		panic("number of arguments must be greater than 0")
	}

	return max(vals...)
}

func min2(vals ...float64) float64 {
	if len(vals) < 1 {
		panic("number of arguments must be greater than 0")
	}

	return min(vals...)
}

func main() {
	fmt.Println(max())           // "0"
	fmt.Println(min())           // "0"
	fmt.Println(max(3))          // "3"
	fmt.Println(min(3))          // "3"
	fmt.Println(max(1, 2, 3, 4)) // "4"
	fmt.Println(min(1, 2, 3, 4)) // "1"

	values := []float64{1, 2, 3, 4}
	fmt.Println(max(values...)) // "4"
	fmt.Println(min(values...)) // "1"

	fmt.Println(max2()) // panic
	fmt.Println(min2()) // panic
}
