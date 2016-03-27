package github

import "time"

const ReposURL = "https://api.github.com/repos/"

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	State     string
	Title     string
	Body      string
	User      *User
	Labels    *[]Label
	Assignee  *Assignee
	Milestone *Milestone
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Label struct {
	Name string
}

type Assignee struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Milestone struct {
	Number  int
	HTMLURL string `json:"html_url"`
	State   string
	Title   string
	DueOn   time.Time `json:"due_on"`
}
