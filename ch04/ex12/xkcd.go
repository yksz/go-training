package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const indexDir = "index" + string(os.PathSeparator)

type comic struct {
	Num        int
	Title      string
	Img        string
	Transcript string
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("usage: %s <search word>\n", os.Args[0])
		os.Exit(0)
	}
	err := searchComics(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

func searchComics(words []string) error {
	files, err := ioutil.ReadDir(indexDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		file, err := os.Open(indexDir + file.Name())
		if err != nil {
			return err
		}
		defer file.Close()
		var comic comic
		if err := json.NewDecoder(file).Decode(&comic); err != nil {
			return err
		}
		found := false
		for _, word := range words {
			found = found || strings.Contains(strings.ToLower(comic.Title), word)
		}
		if found {
			fmt.Println()
			fmt.Printf("#%d URL: \n%s\n", comic.Num, comic.Img)
			fmt.Printf("#%d Transcript: \n%s\n", comic.Num, comic.Transcript)
		}
	}
	return nil
}
