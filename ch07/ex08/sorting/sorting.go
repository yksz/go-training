package sorting

import (
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type CustomSort struct {
	t    []*Track
	keys []string
}

func (x *CustomSort) Sort(t []*Track, key string) {
	x.t = t

	// remove key from x.keys
	newKeys := []string{}
	for _, v := range x.keys {
		if v != key {
			newKeys = append(newKeys, v)
		}
	}
	// append key to the head
	x.keys = append([]string{key}, newKeys...)

	sort.Sort(x)
}

func (x *CustomSort) Len() int {
	return len(x.t)
}

func (x *CustomSort) Less(i, j int) bool {
	for _, key := range x.keys {
		switch key {
		case "title":
			if x.t[i].Title != x.t[j].Title {
				return x.t[i].Title < x.t[j].Title
			}
		case "artist":
			if x.t[i].Artist != x.t[j].Artist {
				return x.t[i].Artist < x.t[j].Artist
			}
		case "album":
			if x.t[i].Album != x.t[j].Album {
				return x.t[i].Album < x.t[j].Album
			}
		case "year":
			if x.t[i].Year != x.t[j].Year {
				return x.t[i].Year < x.t[j].Year
			}
		case "length":
			if x.t[i].Length != x.t[j].Length {
				return x.t[i].Length < x.t[j].Length
			}
		}
	}
	return false
}

func (x *CustomSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}
