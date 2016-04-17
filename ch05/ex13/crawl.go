// Crawler crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	uri "net/url"
	"os"
	"path"

	"golang.org/x/net/html"
)

const rootDir = "./downloads"

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item *uri.URL) []*uri.URL, worklist []*uri.URL) {
	seen := make(map[*uri.URL]bool)
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

func crawl(url *uri.URL) []*uri.URL {
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url *uri.URL) ([]*uri.URL, error) {
	resp, local, n, err := fetch(url)
	if err != nil {
		return nil, fmt.Errorf("fetching %s: %v", url, err)
	}
	if local == "" {
		return nil, nil
	}
	fmt.Printf("%s => %s (%d bytes).\n", url, local, n)

	file, err := os.Open(local)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(file)
	file.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []*uri.URL
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				if link.Host != url.Host {
					continue
				}
				links = append(links, link)
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url *uri.URL) (r *http.Response, filename string, n int64, err error) {
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, "", 0, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	dir, file := path.Split(url.Path)
	if dir == "" {
		dir = "/"
	}
	if file == "" {
		file = "index.html"
	}
	local := rootDir + dir + file

	os.MkdirAll(rootDir+dir, 0755)
	if exists(local) {
		return nil, "", 0, nil
	}
	f, err := os.Create(local)
	if err != nil {
		return nil, "", 0, err
	}

	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return resp, local, n, err
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	var urls []*uri.URL
	for _, s := range os.Args[1:] {
		url, err := uri.Parse(s)
		if err != nil {
			continue
		}
		urls = append(urls, url)
	}
	breadthFirst(crawl, urls)
}
