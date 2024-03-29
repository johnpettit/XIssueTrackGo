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

//GetIssues get all issues
func GetIssues(w http.ResponseWriter, r *http.Request) {
	log.Print("GetIssues called")
	issues := business.GetIssues()
	resp, _ := json.Marshal(issues)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

//GetIssue gets 1 Issue by ID
func GetIssue(w http.ResponseWriter, r *http.Request) {
	log.Print("GetIssue called")
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
	log.Print("CreateIssue called")
	var newissue model.Issue
	json.NewDecoder(r.Body).Decode(&newissue)
	log.Print("LoggedInUserID in create:" + strconv.Itoa(LoggedInUserID))
	newissue, err := business.CreateIssue(newissue, LoggedInUserID)

	if err != nil {
		log.Print("Error Creating Issue")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, _ := json.Marshal(newissue)
	w.Write(resp)
}

//EditIssue edits a Issue
func EditIssue(w http.ResponseWriter, r *http.Request) {
	log.Print("EditIssue called")
	var issue model.Issue
	params := mux.Vars(r)
	issueid, _ := strconv.Atoi(params["issueid"])
	json.NewDecoder(r.Body).Decode(&issue)
	issue, err := business.EditIssue(issueid, issue)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp, _ := json.Marshal(issue)
	w.Write(resp)
}

//DeleteIssue deletes 1 Issue
func DeleteIssue(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteIssue called")
	params := mux.Vars(r)
	issueid, _ := strconv.Atoi(params["issueid"])
	err := business.DeleteIssue(issueid)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
