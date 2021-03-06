package model

import "time"

//User represents a system User
type User struct {
	UserID     int       `json:"userid"`
	FirstName  string    `json:"firstname"`
	LastName   string    `json:"lastname"`
	Email      string    `json:"email"`
	Password   string    `json:"password,omitempty"`
	CreateDate time.Time `json:"createdate"`
	UpdateDate time.Time `json:"updatedate"`
}
