package src

import (
	"fmt"
	jira "github.com/andygrunwald/go-jira"
)

func NewIssue() {

}

func GetIssue(ID string) {
	jiraClient, err := jira.NewClient(nil, "https://jira.infra.online.net/secure/Dashboard.jspa")
	if err != nil {
		panic(err)
	}
	issue, resp, err := jiraClient.Issue.Get(ID, nil)
	if resp.StatusCode > 299 {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
	fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)
}
