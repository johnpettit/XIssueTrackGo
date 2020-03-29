package restapi

import (
	"XIssueTrackGo/business"
	"fmt"
	"net/http"
)

//GetIssue get 1 issue by ID
func GetIssue(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Issue Data")
	new := business.GetIssue("123")
	fmt.Fprintf(w, new.Title)
}
