package links

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"golang.org/x/net/html"
)

var archiveDir = "./archive/"

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(urlStr string) ([]string, error) {
	url, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", urlStr, err)
	}

	filename, n, err := fetch(url)
	if err != nil {
		return nil, fmt.Errorf("fetching %s: %v", url, err)
	}
	if filename == "" {
		return nil, nil
	}
	fmt.Printf("%s => %s (%d bytes).\n", url, filename, n)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	doc, err := html.Parse(file)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := url.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				if link.Host != url.Host {
					log.Printf("different domain: %s\n", link.Host)
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url *url.URL) (filename string, n int64, err error) {
	resp, err := http.Get(url.String())
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	dir, file := path.Split(url.Path)
	if dir == "" {
		dir = "/"
	}
	if file == "" {
		file = "index.html"
	}
	localDir := archiveDir + url.Host + dir
	localFile := localDir + file

	os.MkdirAll(localDir, 0755)
	if exists(localFile) {
		log.Printf("local file already exists: %s\n", localFile)
		return "", 0, nil
	}
	f, err := os.Create(localFile)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return localFile, n, err
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
