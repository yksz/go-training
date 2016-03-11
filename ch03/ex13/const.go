package main

import (
	"fmt"
)

const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000
	YB = ZB * 1000
)

func main() {
	fmt.Printf("KB = %d\n", KB)
	fmt.Printf("MB = %d\n", MB)
	fmt.Printf("GB = %d\n", GB)
	fmt.Printf("TB = %d\n", TB)
	fmt.Printf("PB = %d\n", PB)
	fmt.Printf("EB = %d\n", EB)
	//	fmt.Printf("ZB = %d\n", ZB)
	//	fmt.Printf("YB = %d\n", YB)
}
