package restapi

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var apiPrefix = "/api/v1.0/"

type apiServer struct {
	router *mux.Router
}

//Start entry point
func Start() {
	adapter := mux.NewRouter()

	adapter.HandleFunc("/", heartbeat).Methods("GET")
	adapter.HandleFunc(apiPrefix+"/", heartbeat).Methods("GET")

	adapter.HandleFunc(apiPrefix+"login", Login).Methods("POST")

	//User endpoints
	adapter.HandleFunc(apiPrefix+"user", GetUsers).Methods("GET")
	adapter.HandleFunc(apiPrefix+"user/{userid}", GetUser).Methods("GET")
	adapter.HandleFunc(apiPrefix+"user", CreateUser).Methods("POST")
	adapter.HandleFunc(apiPrefix+"user/{userid}", EditUser).Methods("PUT")
	adapter.HandleFunc(apiPrefix+"user/{userid}", DeleteUser).Methods("DELETE")

	//Issue endpoints
	adapter.HandleFunc(apiPrefix+"issue", GetIssues).Methods("GET")
	adapter.HandleFunc(apiPrefix+"issue/{issueid}", GetIssue).Methods("GET")
	adapter.HandleFunc(apiPrefix+"issue", CreateIssue).Methods("POST")
	adapter.HandleFunc(apiPrefix+"issue/{issueid}", EditIssue).Methods("PUT")
	adapter.HandleFunc(apiPrefix+"issue/{issueid}", DeleteIssue).Methods("DELETE")

	//apiSrv := &http.Server{Addr: "0.0.0.0:8088", Handler: context.ClearHandler(&apiServer{adapter})}
	//log.Fatal(apiSrv.ListenAndServe())
	allowed := handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))
	//log.Fatal(http.ListenAndServe(":8088", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(adapter)))
	log.Fatal(http.ListenAndServe(":8088", allowed(adapter)))
}

func (server *apiServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	//if origin := req.Header.Get("Origin"); origin != "" {
	rw.Header().Set("Access-Contol-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	rw.Header().Set("Content-Type", "application/json")
	//	}

	if req.Method == "OPTIONS" {
		return
	}

	//token checking here
	requri := req.RequestURI
	if (requri != (apiPrefix + "login")) && (requri != "/") && (requri != "/metrics") {
		s := strings.Split(req.Header.Get("Authorization"), " ")
		if len(s) != 2 {
			log.Print("BadAuthHeader")
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenhash := s[1]
		valid, userid := CheckToken(tokenhash)

		req.Header.Add("userid", strconv.Itoa(userid))
		if !valid {
			log.Print("Bad Token")
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	server.router.ServeHTTP(rw, req)
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	log.Print("Home Page")
}
