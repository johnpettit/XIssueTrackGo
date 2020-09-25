package business

import (
	"github.com/johnpettit/XIssueTrackGo/database"
	"github.com/johnpettit/XIssueTrackGo/model"
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

	result := database.DBSession.QueryRow("SELECT id, firstname, lastname, email FROM users WHERE email = ? and password = ?", auth.Email, auth.Password)

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
func CheckToken(tokenhash string) (bool, int) {
	var token model.TokenInternal

	result := database.DBSession.QueryRow("SELECT tokenhash, userid, lasttouch FROM tokens WHERE tokenhash = ?", tokenhash)

	switch err := result.Scan(&token.TokenHash, &token.UserID, &token.LastTouch); err {
	case sql.ErrNoRows:
		//no match    bad token
		return false, -1
	case nil:
		return true, 1
	default:
		//WHAT??
		panic(err)
	}
}

func generateToken(userid int) string {
	var token model.Token
	hasher := md5.New()
	hasher.Write([]byte(strconv.Itoa(userid) + time.Now().Format(time.RFC3339Nano)))
	token.Token = hex.EncodeToString(hasher.Sum(nil))

	del, err := database.DBSession.Prepare("DELETE FROM tokens WHERE userid = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = del.Exec(userid)
	if err != nil {
		panic(err.Error())
	}

	ins, err2 := database.DBSession.Prepare("INSERT INTO tokens (tokenhash, userid) VALUES (?,?)")
	if err2 != nil {
		panic(err2.Error())
	}
	_, err2 = ins.Exec(token.Token, userid)
	if err2 != nil {
		panic(err2.Error())
	}

	return token.Token
}
