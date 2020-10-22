package restapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/johnpettit/XIssueTrackGo/business"
	"github.com/johnpettit/XIssueTrackGo/model"
)

//Login log in a User
func Login(w http.ResponseWriter, r *http.Request) {
	log.Print("Login called")

	var authreq model.AuthRequest
	json.NewDecoder(r.Body).Decode(&authreq)

	token, err := business.UserLogin(authreq)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	log.Print("Logged In!")
	resp, _ := json.Marshal(token)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

//CheckToken checks a tokens validity
func CheckToken(tokenhash string) (bool, int) {
	log.Print("REST - CheckToken called")
	return business.CheckToken(tokenhash)
}
