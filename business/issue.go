package business

import (
	"XIssueTrackGo/database"
	"XIssueTrackGo/model"
	"log"
	"time"
)

//GetIssues gets all Users
func GetIssues() []model.Issue {
	var issue model.Issue
	var issues []model.Issue

	query := "SELECT id, title FROM issues"

	results, err := database.DBSession.Query(query)

	if err != nil {
		log.Print(err)
		return issues
	}

	for results.Next() {
		err = results.Scan(&issue.IssueID, &issue.Title)
		if err != nil {
			log.Print(err)
			return issues
		}
		issues = append(issues, issue)
	}
	return issues
}

//GetIssue gets 1 issue based on ID
func GetIssue(issueID string) *model.Issue {
	new := model.Issue{}
	new.IssueID = "456"
	new.Title = "Hey"
	new.CreateDate = time.Now()
	new.CreatedByUserID = "123"

	return &new
}
