package routing

import "fmt"
import "net/http"

import "github.com/gorilla/mux"

func CreateDiagnosticsRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/healthz", healthzHandler()).Methods(http.MethodGet)
	router.HandleFunc("/readyz", readyzHandler()).Methods(http.MethodGet)
	return router
}

func healthzHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusText(http.StatusOK))
	}
}

func readyzHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusText(http.StatusOK))
	}
}
