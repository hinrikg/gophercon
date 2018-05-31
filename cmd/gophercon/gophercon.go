package main

import "net/http"
import "log"

import "github.com/hinrikg/gophercon/pkg/routing"

func main() {
	log.Printf("Service starting ...")
	router := routing.CreateRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
