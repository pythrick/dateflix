package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	healthCheckPath = "/health-check"
)

func NewHTTPServer() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(healthCheckPath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Alive")
	}).Methods("GET")
	return r
}
