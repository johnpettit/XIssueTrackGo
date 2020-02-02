package issue

import "time"

type Issue struct {
  IssueID string
  Title string
  CreateDate time.Time
  CreatedByUserID string
}

func GetOneIssue(issueID string) *Issue {
  new := Issue{}
  new.IssueID = "456"
  new.Title = "Hey"
  new.CreateDate = time.Now()
  new.CreatedByUserID = "123"

  return &new
}
