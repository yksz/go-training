package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"

	. "./sorting"
)

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func copyTracks(src []*Track) []*Track {
	var dst []*Track
	for _, v := range src {
		t := *v
		dst = append(dst, &t)
	}
	return dst
}

func main() {
	{
		s := new(CustomSort)

		fmt.Println("\nCustom: byYear:")
		t1 := copyTracks(tracks)
		s.Sort(t1, "year")
		printTracks(t1)

		fmt.Println("\nCustom: byTitle:")
		t2 := copyTracks(tracks)
		s.Sort(t2, "title")
		printTracks(t2)
	}
	{
		fmt.Println("\nStable:")
		t := copyTracks(tracks)
		sort.Stable(byYear(t))
		sort.Stable(byTitle(t))
		printTracks(t)
	}
}
