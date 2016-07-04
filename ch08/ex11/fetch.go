// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var done = make(chan struct{})

func mirroredQuery(urls []string) string {
	defer close(done)
	response := make(chan string, len(urls))
	for _, v := range urls {
		go func(url string) {
			response <- request(url)
		}(v)
	}
	return <-response // return the quickest response
}

func request(url string) (response string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Sprintf("fetch: %v\n", err)
	}
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Sprintf("fetch: %v\n", err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("fetch: reading %s: %v\n", url, err)
	}
	return string(b)
}

func main() {
	fmt.Print(mirroredQuery(os.Args[1:]))
}
