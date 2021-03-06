package business

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/johnpettit/XIssueTrackGo/database"
	"github.com/johnpettit/XIssueTrackGo/model"
)

//UserLogin login a user
func UserLogin(auth model.AuthRequest) (model.Token, error) {
	var user model.User
	var token model.Token

	//md5 the Password
	hash := md5.Sum([]byte(auth.Password))
	md5Password := hex.EncodeToString(hash[:])

	result := database.DBSession.QueryRow("SELECT id, firstname, lastname, email FROM users WHERE email = ? and password = ?", auth.Email, md5Password)

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

//CheckToken see if token is valid upates its lasttouched
//TODO add in token expiration
func CheckToken(tokenhash string) (bool, int) {
	log.Print("CheckToken called")
	var token model.TokenInternal

	result := database.DBSession.QueryRow("SELECT tokenhash, userid, lasttouch FROM tokens WHERE tokenhash = ?", tokenhash)

	switch err := result.Scan(&token.TokenHash, &token.UserID, &token.LastTouch); err {
	case sql.ErrNoRows:
		//no match    bad token
		return false, -1
	case nil:
		if updateToken(tokenhash) {
			return true, token.UserID
		}
		return false, -1
	default:
		//WHAT??
		panic(err)
	}
}

//generates a token and inserts into DB
func generateToken(userid int) string {
	log.Print("generateToken called")
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

//updates tokens lasttouched    Should also kill expired tokens??
func updateToken(tokenhash string) bool {
	log.Print("updateToken called")
	up, err := database.DBSession.Prepare("UPDATE tokens SET lasttouch = NOW()  WHERE tokenhash = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = up.Exec(tokenhash)
	if err != nil {
		log.Print(err.Error())
		return false
	}
	return true
}
