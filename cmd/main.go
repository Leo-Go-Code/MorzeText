package main

import (
	"log"
	"os"
	"server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)
	srv := server.New(logger)

	logger.Println("Starting server on :8080")
	err := srv.httpServer.ListenAndServe()
	if err != nil {
		logger.Fatalf("Server failed: %v", err)
	}
}
