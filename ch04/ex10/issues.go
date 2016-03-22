// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	lessThanOneMonthOld := make([]*github.Issue, 0)
	lessThanOneYearOld := make([]*github.Issue, 0)
	moreThanOneYearOld := make([]*github.Issue, 0)
	oneMonthAgo := time.Now().Add(-time.Hour * 24 * 30)
	oneYearAgo := time.Now().Add(-time.Hour * 24 * 365)
	for _, item := range result.Items {
		if item.CreatedAt.After(oneMonthAgo) {
			lessThanOneMonthOld = append(lessThanOneMonthOld, item)
		} else if item.CreatedAt.After(oneYearAgo) {
			lessThanOneYearOld = append(lessThanOneYearOld, item)
		} else {
			moreThanOneYearOld = append(moreThanOneYearOld, item)
		}
	}
	fmt.Printf("\nLess than a month old:\n")
	printIssues(lessThanOneMonthOld)
	fmt.Printf("\nLess than a year old:\n")
	printIssues(lessThanOneYearOld)
	fmt.Printf("\nMore than a year old:\n")
	printIssues(moreThanOneYearOld)
}

func printIssues(issues []*github.Issue) {
	for _, issue := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}
}
