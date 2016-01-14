package main

import (
	"./tempconv"
	"fmt"
)

func main() {
	fmt.Println("AbsoluteZeroK:", tempconv.CToK(tempconv.AbsoluteZeroC))
	fmt.Println("FreezingK:", tempconv.CToK(tempconv.FreezingC))
	fmt.Println("BoilingK:", tempconv.CToK(tempconv.BoilingC))
}
