package business

import "XIssueTrackGo/model"

//GetUser gets 1 User based on ID
func GetUser(userID string) *model.User {
	new := model.User{}
	new.UserID = "123"
	new.FirstName = "Jojo"
	new.Email = "scrap@gmail.com"

	return &new
}

//CreateUser creates a new User
func CreateUser(fName string, lName string, email string) *model.User {
	new := model.User{}
	new.FirstName = fName
	new.LastName = lName
	new.Email = email

	return &new
}