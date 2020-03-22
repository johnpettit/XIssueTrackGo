package main

import (
	"XIssueTrackGo/issue"
	"XIssueTrackGo/user"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	/*
		db, err := sql.Open("mysql", "root:chuck111@/gotest")

		if err != nil {
			fmt.Printf(err.Error())
		}

		defer db.Close()
	*/
	router := mux.NewRouter()

	router.HandleFunc("/", homeLink).Methods("GET")

	router.HandleFunc("/login", loginUser).Methods("POST")

	router.HandleFunc("/user", getUser).Methods("GET")
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user", editUser).Methods("PUT")
	router.HandleFunc("/user", deleteUser).Methods("DELETE")

	router.HandleFunc("/issue", getIssue).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login in this")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	newREST := user.UserREST{}
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
	new := user.CreateUser("Jo", "Bob", "jobob@gmail.com")
	fmt.Fprintf(w, new.Email)
}

func editUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Edit User Data")
	new := user.CreateUser("Jo", "Bob", "jobob@gmail.com")
	fmt.Fprintf(w, new.Email)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Data")
	new := user.CreateUser("Jo", "Bob", "jobob@gmail.com")
	fmt.Fprintf(w, new.Email)
}

func getIssue(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Issue Data")
	new := issue.GetOneIssue("123")
	fmt.Fprintf(w, new.Title)
}
