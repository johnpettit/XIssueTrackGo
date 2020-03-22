package business

import (
	"XIssueTrackGo/model"
	"time"
)

//GetIssue gets 1 issue based on ID
func GetIssue(issueID string) *model.Issue {
	new := model.Issue{}
	new.IssueID = "456"
	new.Title = "Hey"
	new.CreateDate = time.Now()
	new.CreatedByUserID = "123"

	return &new
}
