package main

import (
	"fmt"
)

func square(x int) (result int) {
	defer func() {
		if p := recover(); p != nil {
			result = x * x
		}
	}()
	panic("")
}

func main() {
	fmt.Println(square(16)) // 256
}
