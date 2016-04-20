// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"sort"
)

const (
	temporary = iota + 1
	permanent
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]int)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if seen[item] == temporary {
				fmt.Printf("cycles: %s\n", item)
			} else if _, ok := seen[item]; !ok {
				seen[item] = temporary
				visitAll(m[item])
				seen[item] = permanent
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}
