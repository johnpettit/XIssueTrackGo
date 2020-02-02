package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", homeLink).Methods("GET")

    router.HandleFunc("/login", loginUser).Methods("POST")

		    router.HandleFunc("/user", getUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login in this")
}

func getUser(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Get User Data")
}
