package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

func init() {
	flag.Parse()
}

type player struct {
	name  string
	court chan int
}

func (p *player) rally(to *player, abort <-chan struct{}) {
	count := 0
	for {
		select {
		case ball := <-p.court:
			if *vFlag {
				fmt.Printf("%s: %d\n", p.name, ball)
			}
			count = ball
			ball++
			to.court <- ball
		case <-abort:
			fmt.Printf("%s: rally=%d\n", p.name, count)
			return
		}
	}
}

func newPlayer(name string) *player {
	return &player{name: name, court: make(chan int)}
}

func main() {
	abort := make(chan struct{})
	player1 := newPlayer("player1")
	player2 := newPlayer("player2")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		player1.rally(player2, abort)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		player2.rally(player1, abort)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		select {
		case <-time.After(1 * time.Second):
			abort <- struct{}{}
		}
		wg.Done()
	}()
	player1.court <- 0
	wg.Wait()
}
