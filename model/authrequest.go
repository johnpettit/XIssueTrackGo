package model

//AuthRequest is the request to login
type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
