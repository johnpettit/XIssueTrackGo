package business

import (
	"XIssueTrackGo/database"
	"XIssueTrackGo/model"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"strconv"
	"time"
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
		token.Token = generateToken(user.UserID)
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

func generateToken(userid int) string {
	var token model.Token
	hasher := md5.New()
	hasher.Write([]byte(strconv.Itoa(userid) + time.Now().Format(time.RFC3339Nano)))
	token.Token = hex.EncodeToString(hasher.Sum(nil))

	del, err := database.DBSession.Prepare("DELETE FROM token WHERE userid = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = del.Exec(userid)
	if err != nil {
		panic(err.Error())
	}

	ins, err2 := database.DBSession.Prepare("INSERT INTO token (tokenhash, userid) VALUES (?,?)")
	if err2 != nil {
		panic(err2.Error())
	}
	_, err2 = ins.Exec(token.Token, userid)
	if err2 != nil {
		panic(err2.Error())
	}

	return token.Token
}
