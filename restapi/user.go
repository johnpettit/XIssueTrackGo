package restapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/johnpettit/XIssueTrackGo/business"
	"github.com/johnpettit/XIssueTrackGo/model"

	"github.com/gorilla/mux"
)

//GetUsers gets all Users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Print("GetUsers called")
	users := business.GetUsers()
	resp, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

//GetUser gets 1 User by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	log.Print("GetUser called")
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
	log.Print("CreateUser called")
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

//EditUser edits a User
func EditUser(w http.ResponseWriter, r *http.Request) {
	log.Print("EditUser called")
	var user model.User
	params := mux.Vars(r)
	userid, _ := strconv.Atoi(params["userid"])
	json.NewDecoder(r.Body).Decode(&user)
	user, err := business.EditUser(userid, user)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp, _ := json.Marshal(user)
	w.Write(resp)
}

//DeleteUser deletes 1 User
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteUser called")
	params := mux.Vars(r)
	userid, _ := strconv.Atoi(params["userid"])
	err := business.DeleteUser(userid)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
