package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const cacheDir = "cache" + string(os.PathSeparator)

func init() {
	os.Mkdir(cacheDir, os.ModePerm)
}

func GetIssues(repo string) ([]*Issue, error) {
	filename := cacheDir + strings.Replace(repo, "/", "_", -1)
	if !Exists(filename) {
		url := ReposURL + repo + "/issues"
		if err := fetch(url, filename); err != nil {
			return nil, err
		}
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []*Issue
	if err := json.NewDecoder(file).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func Exists(filename string) bool {
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
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status: %s", resp.Status)
	}

	if _, err := io.Copy(file, resp.Body); err != nil {
		return err
	}
	return nil
}
