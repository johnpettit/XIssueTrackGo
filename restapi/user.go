package restapi

import (
	"XIssueTrackGo/business"
	"XIssueTrackGo/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	params := mux.Vars(r)
	userid, _ := strconv.Atoi(params["userid"])
	user, err := business.GetUser(userid)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp, _ := json.Marshal(user)
	w.Write(resp)
}

//CreateUser create 1 User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newuser model.User

	json.NewDecoder(r.Body).Decode(&newuser)

	newuser, err := business.CreateUser(newuser)

	if err != nil {
		log.Print("Error Creating User")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, _ := json.Marshal(newuser)
	w.Write(resp)
}

/*
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
*/
