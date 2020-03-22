package model

import "time"

//Issue is the Issues logged to the system
type Issue struct {
	IssueID         string
	Title           string
	CreateDate      time.Time
	CreatedByUserID string
}
