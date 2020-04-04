package business

import (
	"XIssueTrackGo/database"
	"XIssueTrackGo/model"
	"database/sql"
	"errors"
)

//UserLogin login a user
func UserLogin(auth model.AuthRequest) (model.Token, error) {
	var user model.User
	var token model.Token

	result := database.DBSession.QueryRow("SELECT id, firstname, lastname, email FROM user WHERE email = ? and password = ?", auth.Email, auth.Password)

	switch err := result.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email); err {
	case sql.ErrNoRows:
		//no match    bad pw?   invalid email?
		return token, errors.New("Email/Password not found")
	case nil:
		token.Token = "37425b789a27ecfa5e6fc59183ac34ca"
		return token, nil
	default:
		//WHAT??
		panic(err)
	}
}

//CheckToken see if token is valid
func CheckToken(tokenhash string) (bool, int, error) {
	return true, 1, nil
}
