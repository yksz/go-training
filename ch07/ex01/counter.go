package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	*c += WordCounter(count)
	return count, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	*c += LineCounter(count)
	return count, nil
}

func main() {
	{
		var c WordCounter
		c.Write([]byte("hello"))
		fmt.Println(c) // "1"

		c = 0 // reset the counter
		var name = "Dolly"
		fmt.Fprintf(&c, "hello, %s", name)
		fmt.Println(c) // "2"
	}

	{
		var c LineCounter
		c.Write([]byte("hello"))
		fmt.Println(c) // "1"

		c = 0 // reset the counter
		var name = "Dolly"
		fmt.Fprintf(&c, "hello, %s\ngoodbye, %s", name, name)
		fmt.Println(c) // "2"
	}
}
