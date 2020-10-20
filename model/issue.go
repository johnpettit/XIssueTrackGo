package model

import "time"

//Issue is the Issues logged to the system
type Issue struct {
	IssueID         int       `json:"issueid"`
	Title           string    `json:"title"`
	CreateDate      time.Time `json:"createdate"`
	UpdateDate      time.Time `json:"updatedate"`
	CreatedByUserID int       `json:"createdbyuserid"`
}
