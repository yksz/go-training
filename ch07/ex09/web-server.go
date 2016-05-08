package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"

	. "../ex08/sorting"
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

func main() {
	customSort := CustomSort{}
	trackTemplate := template.Must(template.ParseFiles("templates/track.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		t := copyTracks(tracks)
		sortKey := r.FormValue("sort")
		if sortKey != "" {
			sort.Sort(customSort.Sort(t, sortKey))
		}
		if err := trackTemplate.Execute(w, t); err != nil {
			log.Print(err)
		}
	})
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js"))))
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func copyTracks(src []*Track) []*Track {
	var dst []*Track
	for _, v := range src {
		t := *v
		dst = append(dst, &t)
	}
	return dst
}
