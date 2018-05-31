package routing

import "net/http"
import "fmt"
import "log"

import "github.com/gorilla/mux"

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", homeHandler()).Methods(http.MethodGet)
	return r
}

func homeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request is processing %s", r.URL.Path)
		fmt.Fprint(w, "Hello!")
	}
}
