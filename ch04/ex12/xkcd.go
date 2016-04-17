package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	prefixURL = "https://xkcd.com/"
	suffixURL = "/info.0.json"
	cacheDir  = "cache" + string(os.PathSeparator)
)

var m sync.Mutex

type comic struct {
	Num        int
	Title      string
	Img        string
	Transcript string
}

func init() {
	os.Mkdir(cacheDir, os.ModePerm)
}

func main() {
	if len(os.Args) <= 2 {
		fmt.Printf("usage: %s <search range> <search word>\n", os.Args[0])
		os.Exit(0)
	}
	from, to, err := getRange(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if err := searchComics(from, to, os.Args[2:]); err != nil {
		log.Fatal(err)
	}
}

func getRange(str string) (int, int, error) {
	s := strings.SplitN(str, "-", -1)
	if len(s) != 2 {
		return 0, 0, fmt.Errorf("invalid range: %s", str)
	}
	from, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, 0, err
	}
	to, err := strconv.Atoi(s[1])
	if err != nil {
		return 0, 0, err
	}
	return from, to, nil
}

func searchComics(from, to int, words []string) error {
	ch := make(chan string)
	for i := from; i <= to; i++ {
		go searchComic(i, words, ch)
	}
	for i := from; i <= to; i++ {
		s := <-ch
		if s != "" {
			log.Print(s)
		}
	}
	return nil
}

func searchComic(num int, words []string, ch chan<- string) {
	filename := cacheDir + strconv.Itoa(num)
	if !exists(filename) {
		url := prefixURL + strconv.Itoa(num) + suffixURL
		if err := fetch(url, filename); err != nil {
			ch <- fmt.Sprint(err)
			return
		}
	}
	file, err := os.Open(filename)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer file.Close()
	var comic comic
	if err := json.NewDecoder(file).Decode(&comic); err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	found := false
	for _, word := range words {
		found = found || strings.Contains(strings.ToLower(comic.Title), word)
	}
	if found {
		m.Lock()
		fmt.Println()
		fmt.Printf("#%d URL: \n%s\n", comic.Num, comic.Img)
		fmt.Printf("#%d Transcript: \n%s\n", comic.Num, comic.Transcript)
		m.Unlock()
	}
	ch <- ""
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
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
