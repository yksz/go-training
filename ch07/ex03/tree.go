package main

import (
	"fmt"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	if t == nil {
		return ""
	}
	s := ""
	s += t.left.String()
	s += strconv.Itoa(t.value) + " "
	s += t.right.String()
	return s
}

func main() {
	values := []int{1, 5, 2, 4, 3}
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	fmt.Println(root) // "1 2 3 4 5"
}
