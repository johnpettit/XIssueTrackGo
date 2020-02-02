package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", homeLink).Methods("GET")

    router.HandleFunc("/login", loginUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login in this")
}


