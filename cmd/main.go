package main

import (
	"MorzeText/internal/server"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)
	srv := server.NewServer(logger)

	logger.Println("Starting server on :8080")
	err := srv.HTTPServer.ListenAndServe()
	if err != nil {
		logger.Fatalf("Server failed: %v", err)
	}
}
