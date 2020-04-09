package model

import "time"

//TokenInternal is the token for a successful login
type TokenInternal struct {
	TokenHash string
	UserID    int
	LastTouch time.Time
}
