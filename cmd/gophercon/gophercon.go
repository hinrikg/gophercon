package main

import "log"
import "os"

import "github.com/hinrikg/gophercon/pkg/routing"
import "github.com/hinrikg/gophercon/pkg/webserver"

func main() {
	log.Printf("Service starting ...")

	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("Service port was not set!")
	}

	router := routing.CreateRouter()
	server := webserver.New("", port, router)
	log.Fatal(server.Start())
}
