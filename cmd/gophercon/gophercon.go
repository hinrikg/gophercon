package main

import "log"
import "os"

import "github.com/hinrikg/gophercon/pkg/routing"
import "github.com/hinrikg/gophercon/pkg/webserver"
import "github.com/hinrikg/gophercon/version"

func main() {
	log.Printf(
		"Service is starting, version is %s, commit is %s, time is %s...",
		version.Release, version.Commit, version.BuildTime,
	)

	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("Service port was not set!")
	}

	router := routing.CreateRouter()
	server := webserver.New("", port, router)
	log.Fatal(server.Start())
}
