package restapi

import (
	"XIssueTrackGo/business"
	"XIssueTrackGo/model"
	"encoding/json"
	"fmt"
	"net/http"
)

//GetUsers gets all Users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := business.GetUsers()
	resp, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

//GetUser gets 1 User by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	newREST := model.User{}
	newREST.UserID = 123
	newREST.FirstName = "JoJo"
	newREST.Email = "jojo@gmail.com"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newREST); err != nil {
		panic(err)
	}
}

//CreateUser create 1 User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User Data")
	new := business.CreateUser("Jo", "Bob", "jobob@gmail.com")
	fmt.Fprintf(w, new.Email)
}

//EditUser edits a User
func EditUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Edit User Data")
	new := business.CreateUser("Jo", "Bob", "jobob@gmail.com")
	fmt.Fprintf(w, new.Email)
}

//DeleteUser deletes 1 User
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Data")
	new := business.CreateUser("Jo", "Bob", "jobob@gmail.com")
	fmt.Fprintf(w, new.Email)
}
