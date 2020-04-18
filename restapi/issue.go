package restapi

import (
	"XIssueTrackGo/business"
	"encoding/json"
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
