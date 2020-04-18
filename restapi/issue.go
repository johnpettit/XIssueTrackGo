package restapi

import (
	"XIssueTrackGo/business"
	"encoding/json"
	"net/http"
)

//GetIssues get all issues
func GetIssues(w http.ResponseWriter, r *http.Request) {
	issues := business.GetIssues()
	resp, _ := json.Marshal(issues)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
