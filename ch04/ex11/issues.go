package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	apiURL   = "https://api.github.com/repos/"
	required = "REQUIRED"
)

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func getIssues(repo string) error {
	resp, err := http.Get(apiURL + repo + "/issues?state=all")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status: %s", resp.Status)
	}
	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return err
	}
	for _, issue := range issues {
		fmt.Printf("#%-5d %-6s %9.9s %.55s\n",
			issue.Number, issue.State, issue.User.Login, issue.Title)
	}
	return nil
}

func getIssue(repo string, num int) error {
	resp, err := http.Get(apiURL + repo + "/issues/" + strconv.Itoa(num))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status: %s", resp.Status)
	}
	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return err
	}
	fmt.Printf("#%d: %s\n", issue.Number, strings.Title(issue.State))
	fmt.Printf("Title: %s\n", issue.Title)
	fmt.Printf("User: %s\n", issue.User.Login)
	fmt.Printf("Comment: %s\n", issue.Body)
	return nil
}

func createIssue(repo string, user, passwd, title, comment string) error {
	url := apiURL + repo + "/issues"
	body := fmt.Sprintf("{\"title\":\"%s\",\"body\":\"%s\"}", title, comment)
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	req.SetBasicAuth(user, passwd)
	req.ContentLength = int64(len(body))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status: %s", resp.Status)
	}
	return nil
}

func editIssue(repo string, num int, user, passwd, title, comment string) error {
	url := apiURL + repo + "/issues/" + strconv.Itoa(num)
	body := fmt.Sprintf("{\"title\":\"%s\",\"body\":\"%s\"}", title, comment)
	req, err := http.NewRequest("PATCH", url, strings.NewReader(body))
	req.SetBasicAuth(user, passwd)
	req.ContentLength = int64(len(body))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status: %s", resp.Status)
	}
	return nil
}

func closeIssues(repo string, num int, user, passwd string) error {
	url := apiURL + repo + "/issues/" + strconv.Itoa(num)
	body := "{\"state\":\"closed\"}"
	req, err := http.NewRequest("PATCH", url, strings.NewReader(body))
	req.SetBasicAuth(user, passwd)
	req.ContentLength = int64(len(body))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status: %s", resp.Status)
	}
	return nil
}

func main() {
	var create, edit, close bool
	flag.BoolVar(&create, "create", false, "mode: create")
	flag.BoolVar(&edit, "edit", false, "mode: edit")
	flag.BoolVar(&close, "close", false, "mode: close")
	id := flag.String("id", required, "issue: id: e.g. golang/go or golang/go#1")
	user := flag.String("user", "", "auth: username")
	title := flag.String("title", "", "issue: title")
	comment := flag.String("comment", "", "issue: comment")
	flag.Parse()

	if *id == required {
		fmt.Fprintf(os.Stderr, "id is %s\n\n", required)
		flag.PrintDefaults()
		os.Exit(1)
	}
	repo, num := parseIssueID(*id)

	passwd := ""
	if *user != "" {
		passwd = getPassword()
	}

	var err error
	if create {
		err = createIssue(repo, *user, passwd, *title, *comment)
	} else if edit {
		err = editIssue(repo, num, *user, passwd, *title, *comment)
	} else if close {
		err = closeIssues(repo, num, *user, passwd)
	} else {
		if num > 0 {
			err = getIssue(repo, num)
		} else {
			err = getIssues(repo)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}

func parseIssueID(id string) (string, int) {
	s := strings.SplitN(id, "#", 2)
	if len(s) == 2 {
		num, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		return s[0], num
	}
	return s[0], 0
}

func getPassword() string {
	fmt.Printf("Password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		log.Fatal(err)
	}
	return string(password)
}
