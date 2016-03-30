package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
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
		filename := strings.Replace(title, " ", "_", -1) + ".jpg"
		if err := fetch(result.Poster, filename); err != nil {
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

func fetch(url, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		return err
	}
	return nil
}
