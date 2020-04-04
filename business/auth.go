package business

import (
	"XIssueTrackGo/model"
	"errors"
)

//UserLogin login a user
func UserLogin(auth model.AuthRequest) (model.Token, error) {
	var token model.Token
	if auth.Email == "john.pettit@cox.net" {
		token.Token = "37425b789a27ecfa5e6fc59183ac34ca" //md5 of chuck111
		return token, nil
	}
	return token, errors.New("No Match")
}

//CheckToken see if token is valid
func CheckToken(tokenhash string) (bool, int, error) {
	return true, 1, nil
}
