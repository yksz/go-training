package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sort"
	"strings"
)

type clock struct {
	location string
	time     string
}

type byLocation []*clock

func (t byLocation) Len() int           { return len(t) }
func (t byLocation) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t byLocation) Less(i, j int) bool { return t[i].location < t[j].location }

func main() {
	ch := make(chan *clock)
	nconn := 0
	for _, arg := range os.Args[1:] {
		s := strings.Split(arg, "=")
		if len(s) < 2 {
			continue
		}
		location := s[0]
		addr := s[1]
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn, location, ch)
		nconn++
	}
	printClockWall(nconn, ch)
}

func handle(conn net.Conn, location string, out chan<- *clock) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		out <- &clock{location, strings.TrimSpace(line)}
	}
}

func printClockWall(nclocks int, in <-chan *clock) {
	var clocks []*clock
	for {
		clock := <-in
		clocks = append(clocks, clock)
		if len(clocks) == nclocks {
			sort.Sort(byLocation(clocks))
			fmt.Printf("\r")
			for _, c := range clocks {
				fmt.Printf("%s: %s\t", c.location, c.time)
			}
			clocks = nil // clear
		}
	}
}
