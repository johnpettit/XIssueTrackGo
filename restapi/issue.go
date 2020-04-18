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

//GetIssues get all issues
func GetIssues(w http.ResponseWriter, r *http.Request) {
	issues := business.GetIssues()
	resp, _ := json.Marshal(issues)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

//GetIssue gets 1 Issue by ID
func GetIssue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	issueid, _ := strconv.Atoi(params["issueid"])
	issue, err := business.GetIssue(issueid)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp, _ := json.Marshal(issue)
	w.Write(resp)
}

//CreateIssue create 1 Issue
func CreateIssue(w http.ResponseWriter, r *http.Request) {
	var newissue model.Issue
	json.NewDecoder(r.Body).Decode(&newissue)
	newissue, err := business.CreateIssue(newissue)

	if err != nil {
		log.Print("Error Creating Issue")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, _ := json.Marshal(newissue)
	w.Write(resp)
}
