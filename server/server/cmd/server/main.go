package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	server, err := Compose(8080)
	if err != nil {
		log.Fatalf("Cannot initialize chat server: %s", err)
	} else {
		go func() {
			log.Println("Starting chat server...")

			err := server.Start()
			if err == http.ErrServerClosed {
				log.Printf("HTTP server stopped")
			} else {
				log.Fatalf("Cannot start HTTP server: %s", err)
			}
		}()

		// Wait for Ctrl-C signal.
		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt)
		<-sigChannel

		if err := server.Stop(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error stopping the server: %s", err)
		}
	}
}
