package main

import "log"
import "os"
import "os/signal"
import "syscall"

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

	internalPort := os.Getenv("INTERNAL_PORT")
	if len(internalPort) == 0 {
		log.Fatal("Internal port was not set!")
	}

	shutdown := make(chan error, 2)

	router := routing.CreateRouter()
	server := webserver.New("", port, router)
	go func() {
		error := server.Start()
		shutdown <- error
	}()

	internalRouter := routing.CreateDiagnosticsRouter()
	internalServer := webserver.New("", internalPort, internalRouter)
	go func() {
		error := internalServer.Start()
		shutdown <- error
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case killSignal := <-interrupt:
		log.Printf("Got %s. Stopping ...", killSignal)
	case error := <-shutdown:
		log.Printf("Got an error '%s'. Stopping ...", error)
	}

	log.Print(server.Stop())
	log.Print(internalServer.Stop())
}
