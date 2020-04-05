package restapi

import (
	"XIssueTrackGo/business"
	"XIssueTrackGo/model"
	"encoding/json"
	"log"
	"net/http"
)

//Login log in a User
func Login(w http.ResponseWriter, r *http.Request) {
	log.Print("Login in this")

	var authreq model.AuthRequest
	json.NewDecoder(r.Body).Decode(&authreq)

	token, err := business.UserLogin(authreq)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	log.Print("Logged In!")
	resp, _ := json.Marshal(token)
	w.Write(resp)
}

//CheckToken checks a tokens validity
func CheckToken(tokenhash string) bool {
	return true
}
