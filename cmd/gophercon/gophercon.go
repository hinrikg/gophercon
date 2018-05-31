package main

import "net/http"
import "fmt"
import "log"

import "github.com/gorilla/mux"

func main() {
	log.Printf("Service starting ...")

	r := mux.NewRouter()
	r.HandleFunc("/home", homeHandler()).Methods(http.MethodGet)

	http.ListenAndServe(":8000", r)
}

func homeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request is processing %s", r.URL.Path)
		fmt.Fprint(w, "Hello!")
	}
}
