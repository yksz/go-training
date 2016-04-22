// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]set{
	"algorithms": newSet("data structures"),
	"calculus":   newSet("linear algebra"),

	"compilers": newSet(
		"data structures",
		"formal languages",
		"computer organization",
	),

	"data structures":       newSet("discrete math"),
	"databases":             newSet("data structures"),
	"discrete math":         newSet("intro to programming"),
	"formal languages":      newSet("discrete math"),
	"networks":              newSet("operating systems"),
	"operating systems":     newSet("data structures", "computer organization"),
	"programming languages": newSet("data structures", "computer organization"),
}

type set map[string]bool

func newSet(s ...string) set {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = true
	}
	return m
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]set) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items set)

	visitAll = func(items set) {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	keys := newSet()
	for key := range m {
		keys[key] = true
	}

	visitAll(keys)
	return order
}
