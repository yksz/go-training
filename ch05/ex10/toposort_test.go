package main

import "testing"

func TestTopoSort(t *testing.T) {
	sortItems := reverse(topoSort(prereqs))
	for key, set := range prereqs {
		for i, item := range sortItems {
			if item == key {
				if !containsAll(sortItems[i:], set) {
					t.Errorf("%s", key)
				}
			}
		}
	}
}

func containsAll(s []string, vals set) bool {
	base := newSet(s...)
	for val := range vals {
		if _, ok := base[val]; !ok {
			return false
		}
	}
	return true
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
