// Crawl crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

type link struct {
	url   string
	depth int
}

var depth int

func init() {
	flag.IntVar(&depth, "depth", 1, "depth")
	flag.Parse()
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func newLinks(urls []string, depth int) []link {
	var links []link
	for _, url := range urls {
		links = append(links, link{url: url, depth: depth})
	}
	return links
}

func main() {
	worklist := make(chan []link)  // lists of URLs, may have duplicates
	unseenLinks := make(chan link) // de-duplicated URLs

	defer func() {
		close(worklist)
		close(unseenLinks)
	}()

	// Add command-line arguments to worklist.
	go func() { worklist <- newLinks(flag.Args(), 0) }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link.url)
				go func() { worklist <- newLinks(foundLinks, link.depth+1) }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if link.depth > depth {
				continue
			}
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}
