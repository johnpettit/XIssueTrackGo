package main

import (
	"XIssueTrackGo/business"
	"XIssueTrackGo/model"
	"XIssueTrackGo/restapi"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", heartbeat).Methods("GET")

	router.HandleFunc("/login", restapi.Login).Methods("POST")

	router.HandleFunc("/user", getUser).Methods("GET")
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user", editUser).Methods("PUT")
	router.HandleFunc("/user", deleteUser).Methods("DELETE")

	router.HandleFunc("/issue", getIssue).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	log.Print("Home Page")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	newREST := model.User{}
	newREST.UserID = "123"
	newREST.FirstName = "JoJo"
	newREST.Email = "jojo@gmail.com"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newREST); err != nil {
		panic(err)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User Data")
	new := business.CreateUser("Jo", "Bob", "jobob@gmail.com")
	fmt.Fprintf(w, new.Email)
}

func editUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Edit User Data")
	new := business.CreateUser("Jo", "Bob", "jobob@gmail.com")
	fmt.Fprintf(w, new.Email)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Data")
	new := business.CreateUser("Jo", "Bob", "jobob@gmail.com")
	fmt.Fprintf(w, new.Email)
}

func getIssue(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Issue Data")
	new := business.GetIssue("123")
	fmt.Fprintf(w, new.Title)
}
