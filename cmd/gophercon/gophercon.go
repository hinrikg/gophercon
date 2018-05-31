package main

import "net/http"
import "fmt"
import "log"

func main() {
	log.Printf("Service starting ...")

	http.HandleFunc("/home", homeHandler())
	http.ListenAndServe(":8000", nil)
}

func homeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request is processing %s", r.URL.Path)
		fmt.Fprint(w, "Hello!")
	}
}
