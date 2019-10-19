package main

import (
	"log"
	"os"
	"context"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/andrewozhegov/k8s-introduction/version"
	"github.com/andrewozhegov/k8s-introduction/handlers"
)

// DEFAULTPORT returns default port number
const DEFAULTPORT = "8000"

func main() {
	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		port = DEFAULTPORT
	}

	r := handlers.Router(version.RELEASE, version.COMMIT, version.REPO)

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// this channel is for graceful shutdown:
	// if we receive an error, we can send it here to notify the server to be stopped
	shutdown := make(chan struct{}, 1)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			shutdown <- struct{}{}
			log.Printf("%v", err)
		}
	}()
	log.Print("The service is ready to listen and serve.")

	select {
	case killSignal := <-interrupt:
		switch killSignal {
		case os.Interrupt:
			log.Print("Got SIGINT...")
		case syscall.SIGTERM:
			log.Print("Got SIGTERM...")
		}
	case <-shutdown:
		log.Printf("Got an error...")
	}

	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Done")
}
