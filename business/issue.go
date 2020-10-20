package business

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/johnpettit/XIssueTrackGo/database"
	"github.com/johnpettit/XIssueTrackGo/model"
)

//GetIssues gets all Users
func GetIssues() []model.Issue {
	var issue model.Issue
	var issues []model.Issue

	query := "SELECT id, title, createdate, updatedate, createdbyuserid FROM issues"

	results, err := database.DBSession.Query(query)

	if err != nil {
		log.Print(err)
		return issues
	}

	for results.Next() {
		err = results.Scan(&issue.IssueID, &issue.Title, &issue.CreateDate, &issue.UpdateDate, &issue.CreatedByUserID)
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

	query := "SELECT id, title, createdate, updatedate, createdbyuserid FROM issues WHERE id = ?"

	result := database.DBSession.QueryRow(query, values...)

	switch err := result.Scan(&issue.IssueID, &issue.Title, &issue.CreateDate, &issue.UpdateDate, &issue.CreatedByUserID); err {
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
	var createdate = time.Now()
	createdate.Format(time.RFC3339)
	var updatedate = time.Now()
	updatedate.Format(time.RFC3339)
	ins, err := database.DBSession.Prepare("INSERT INTO issues (title, createdate, updatedate, createdbyuserid) VALUES (?,?,?,1)")
	if err != nil {
		log.Print(err)
		return newissue, errors.New("Error Creating Issue")
	}

	insres, err2 := ins.Exec(newissue.Title, createdate, updatedate)
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
	newissue.CreateDate = createdate
	newissue.UpdateDate = updatedate
	newissue.CreatedByUserID = 1
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
