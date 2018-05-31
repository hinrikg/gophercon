package main

import "log"
import "net/http"
import "os"

import "github.com/hinrikg/gophercon/pkg/routing"

func main() {
	log.Printf("Service starting ...")

	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("Service port was not set!")
	}

	router := routing.CreateRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
