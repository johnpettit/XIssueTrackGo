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

//CreateIssue creates a new Issue
func CreateIssue(newissue model.Issue) (model.Issue, error) {
	ins, err := database.DBSession.Prepare("INSERT INTO issues (title) VALUES (?)")
	if err != nil {
		log.Print(err)
		return newissue, errors.New("Error Creating Issue")
	}

	insres, err2 := ins.Exec(newissue.Title)
	if err2 != nil {
		log.Print(err2)
		return newissue, errors.New("Error Creating Issue")
	}

	id, err3 := insres.LastInsertId()
	if err3 != nil {
		log.Print(err3)
		return newissue, errors.New("Error Creating Issue")
	}
	newissue.IssueID = int(id)
	return newissue, nil
}

//EditIssue edits an Issue
func EditIssue(issueid int, issue model.Issue) (model.Issue, error) {
	upd, err := database.DBSession.Prepare("UPDATE issues SET title = ? WHERE id = ?")
	if err != nil {
		log.Print(err)
		return issue, errors.New("Error Editing Issue")
	}

	_, err2 := upd.Exec(issue.Title, issueid)
	if err2 != nil {
		log.Print(err2)
		return issue, errors.New("Error Editing Issue")
	}
	issue, err = GetIssue(issueid)
	return issue, nil
}

//DeleteIssue deletes an Issue
func DeleteIssue(issueid int) error {
	upd, err := database.DBSession.Prepare("DELETE FROM issues WHERE id = ?")
	if err != nil {
		log.Print(err)
		return errors.New("Error Deleting Issue")
	}

	_, err2 := upd.Exec(issueid)
	if err2 != nil {
		log.Print(err2)
		return errors.New("Error Deleting Issue")
	}
	return nil
}
