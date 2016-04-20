package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func tree(dir string) []string {
	fmt.Println(dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Print(err)
		return nil
	}
	var dirs []string
	for _, file := range files {
		name := dir + string(os.PathSeparator) + file.Name()
		if file.IsDir() {
			dirs = append(dirs, name)
		} else {
			fmt.Println(name)
		}
	}
	return dirs
}

func main() {
	breadthFirst(tree, os.Args[1:])
}
