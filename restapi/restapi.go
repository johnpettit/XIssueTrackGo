package restapi

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type apiServer struct {
	router *mux.Router
}

//Start entry point
func Start() {
	adapter := mux.NewRouter()

	recordMetrics()

	adapter.HandleFunc("/", heartbeat).Methods("GET")

	adapter.HandleFunc("/login", Login).Methods("POST")

	adapter.HandleFunc("/user", GetUser).Methods("GET")
	//adapter.HandleFunc("/user", createUser).Methods("POST")
	//adapter.HandleFunc("/user", editUser).Methods("PUT")
	//adapter.HandleFunc("/user", deleteUser).Methods("DELETE")

	adapter.HandleFunc("/issue", GetIssue).Methods("GET")

	//Prometheus metrics
	adapter.Handle("/metrics", promhttp.Handler())

	apiSrv := &http.Server{Addr: "0.0.0.0:8080", Handler: context.ClearHandler(&apiServer{adapter})}
	log.Fatal(apiSrv.ListenAndServe())
}

func (server *apiServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Contol-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	if req.Method == "OPTIONS" {
		return
	}

	//token checking here

	server.router.ServeHTTP(rw, req)
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	log.Print("Home Page")
}

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)
