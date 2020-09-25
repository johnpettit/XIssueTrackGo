package business

import (
	"XIssueTrackGo/database"
	"XIssueTrackGo/model"
	"database/sql"
	"errors"
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
func GetUser(userid int) (model.User, error) {
	var user model.User
	var values []interface{}
	values = append(values, userid)

	query := "SELECT id, firstname, lastname, email FROM users WHERE id = ?"

	result := database.DBSession.QueryRow(query, values...)

	switch err := result.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email); err {
	case sql.ErrNoRows:
		return user, errors.New("No User with that ID")
	case nil:
		return user, nil
	default:
		panic(err)
	}
}

//CreateUser creates a new User
func CreateUser(newuser model.User) (model.User, error) {
	ins, err := database.DBSession.Prepare("INSERT INTO users (firstname, lastname, email, password) VALUES (?,?,?,?)")
	if err != nil {
		log.Print(err)
		return newuser, errors.New("Error Creating User")
	}

	insres, err2 := ins.Exec(newuser.FirstName, newuser.LastName, newuser.Email, newuser.Password)
	if err2 != nil {
		log.Print(err2)
		return newuser, errors.New("Error Creating User")
	}

	id, err3 := insres.LastInsertId()
	if err3 != nil {
		log.Print(err3)
		return newuser, errors.New("Error Creating User")
	}
	newuser.UserID = int(id)
	return newuser, nil
}

//EditUser edits a User
func EditUser(userid int, user model.User) (model.User, error) {
	upd, err := database.DBSession.Prepare("UPDATE users SET firstname = ?, lastname = ? WHERE id = ?")
	if err != nil {
		log.Print(err)
		return user, errors.New("Error Editing User")
	}

	_, err2 := upd.Exec(user.FirstName, user.LastName, userid)
	if err2 != nil {
		log.Print(err2)
		return user, errors.New("Error Editing User")
	}
	user, err = GetUser(userid)
	return user, nil
}

//DeleteUser deletes a User
func DeleteUser(userid int) error {
	upd, err := database.DBSession.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		log.Print(err)
		return errors.New("Error Deleting User")
	}

	_, err2 := upd.Exec(userid)
	if err2 != nil {
		log.Print(err2)
		return errors.New("Error Deleting User")
	}
	return nil
}
