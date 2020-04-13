package business

import (
	"XIssueTrackGo/database"
	"XIssueTrackGo/model"
	"log"
)

//GetUsers gets all Users
func GetUsers() []model.User {
	var user model.User
	var users []model.User

	query := "SELECT id, firstname, lastname, email FROM users"

	results, err := database.DBSession.Query(query)

	if err != nil {
		log.Print(err)
		return users
	}

	for results.Next() {
		err = results.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			log.Print(err)
			return users
		}
		users = append(users, user)
	}
	return users
}

//GetUser gets 1 User based on ID
func GetUser(userid int) model.User {
	new := model.User{}
	new.UserID = userid
	new.FirstName = "Jojo"
	new.Email = "scrap@gmail.com"

	return new
}

//CreateUser creates a new User
func CreateUser(fName string, lName string, email string) *model.User {
	new := model.User{}
	new.FirstName = fName
	new.LastName = lName
	new.Email = email

	return &new
}
