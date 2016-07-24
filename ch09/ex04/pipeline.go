package main

import (
	"flag"
	"fmt"
	"time"
)

const ngoroutine = 1000000

var vFlag = flag.Bool("v", false, "show verbose progress messages")

func init() {
	flag.Parse()
}

func main() {
	start := time.Now()
	pipeline()
	fmt.Printf("Total: %s\n", time.Since(start))
}

func pipeline() {
	in := make(chan int)
	out := make(chan int)
	begin := in
	var end chan int

	for i := 1; i <= ngoroutine; i++ {
		go func(id int, in chan int, out chan int) {
			v := <-in
			if *vFlag {
				fmt.Printf("goroutine_%d: <- %d \n", id, v)
			}
			out <- id
		}(i, in, out)
		end = out
		in = out
		out = make(chan int)
	}

	start := time.Now()
	begin <- 0
	v := <-end
	if *vFlag {
		fmt.Printf("main: <- %d\n", v)
	}
	fmt.Printf("Transit: %s\n", time.Since(start))
}
