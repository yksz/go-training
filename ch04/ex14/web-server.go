package main

import (
	"log"
	"net/http"

	"./github"
)

import "html/template"

var bugReportList = template.Must(template.New("bugreportlist").Parse(`
<h1>Bug Reports</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .}}{{$issue := .}}
{{range .Labels}}
{{if eq .Name "bug"}}
<tr>
  <td><a href='{{$issue.HTMLURL}}'>{{$issue.Number}}</td>
  <td>{{$issue.State}}</td>
  <td><a href='{{$issue.User.HTMLURL}}'>{{$issue.User.Login}}</a></td>
  <td><a href='{{$issue.HTMLURL}}'>{{$issue.Title}}</a></td>
</tr>
{{end}}
{{end}}
{{end}}
</table>
`))

var milestoneList = template.Must(template.New("milestonelist").Parse(`
<h1>Milestones</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>Title</th>
  <th>Due on</th>
</tr>
{{range .}}
{{if .Milestone}}
<tr>
  <td><a href='{{.Milestone.HTMLURL}}'>{{.Milestone.Number}}</td>
  <td>{{.Milestone.State}}</td>
  <td><a href='{{.Milestone.HTMLURL}}'>{{.Milestone.Title}}</a></td>
  <td>{{.Milestone.DueOn.Format "2006-01-02"}}</td>
</tr>
{{end}}
{{end}}
</table>
`))

var userList = template.Must(template.New("userlist").Parse(`
<h1>Users</h1>
<table>
{{range .}}
<tr>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		repo := r.FormValue("repo")
		if repo == "" {
			return
		}

		result, err := github.GetIssues(repo)
		if err != nil {
			log.Print(err)
		}
		if err := bugReportList.Execute(w, result); err != nil {
			log.Print(err)
		}
		m := removeDuplicateMilestones(result)
		if err := milestoneList.Execute(w, m); err != nil {
			log.Print(err)
		}
		u := removeDuplicateUsers(result)
		if err := userList.Execute(w, u); err != nil {
			log.Print(err)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func removeDuplicateMilestones(issues []*github.Issue) []*github.Issue {
	result := make([]*github.Issue, 0)
	for _, v := range issues {
		if v.Milestone == nil {
			continue
		}
		if !contains(result, func(issue *github.Issue) bool {
			return issue.Milestone.Number == v.Milestone.Number
		}) {
			result = append(result, v)
		}
	}
	return result
}

func removeDuplicateUsers(issues []*github.Issue) []*github.Issue {
	result := make([]*github.Issue, 0)
	for _, v := range issues {
		if !contains(result, func(issue *github.Issue) bool {
			return issue.User.Login == v.User.Login
		}) {
			result = append(result, v)
		}
	}
	return result
}

func contains(issues []*github.Issue, predicate func(*github.Issue) bool) bool {
	for _, issue := range issues {
		if predicate(issue) {
			return true
		}
	}
	return false
}
