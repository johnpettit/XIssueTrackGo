package business

import (
	"XIssueTrackGo/database"
	"XIssueTrackGo/model"
	"database/sql"
	"errors"
	"log"
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

//GetIssue gets 1 Issue based on ID
func GetIssue(issueid int) (model.Issue, error) {
	var issue model.Issue
	var values []interface{}
	values = append(values, issueid)

	query := "SELECT id, title FROM issues WHERE id = ?"

	result := database.DBSession.QueryRow(query, values...)

	switch err := result.Scan(&issue.IssueID, &issue.Title); err {
	case sql.ErrNoRows:
		return issue, errors.New("No Issue with that ID")
	case nil:
		return issue, nil
	default:
		panic(err)
	}
}
