// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	filename := getFileName(url)
	file, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer file.Close()

	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s \t-> %s", secs, nbytes, url, filename)
}

func getFileName(url string) string {
	basename := url
	index := strings.Index(basename, "://") + len("://")
	if index > 0 {
		basename = basename[index:len(basename)]
	}

	filename := basename
	for i := 1; i <= 10; i++ {
		if !existsFile(filename) {
			return filename
		}
		filename = basename + "." + strconv.Itoa(i)
	}
	return basename
}

func existsFile(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
