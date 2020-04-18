package model

import "time"

//Issue is the Issues logged to the system
type Issue struct {
	IssueID         int `json:"issueid"`
	Title           string
	CreateDate      time.Time
	CreatedByUserID string
}
