package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

const omdbURL = "https://omdbapi.com/"

type movieSearchResult struct {
	Response string
	Poster   string
	Error    string
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("usage: %s <movie title>\n", os.Args[0])
		os.Exit(0)
	}
	title := os.Args[1]
	downloadMoviePoster(title)
}

func downloadMoviePoster(title string) {
	result, err := obtainMovieInformation(title)
	if err != nil {
		log.Fatal(err)
	}
	if result.Response == "True" {
		err := download(result.Poster, title+".jpg")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(result.Error)
	}
}

func obtainMovieInformation(title string) (*movieSearchResult, error) {
	q := url.QueryEscape(title)
	resp, err := http.Get(omdbURL + "?t=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result movieSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func download(fromURL string, toFile string) error {
	resp, err := http.Get(fromURL)
	if err != nil {
		return err
	}
	file, err := os.Create(toFile)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
