package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(rootDir)))

	port := ":8080"
	log.Printf("Starting server on port %s, serving files from %s", port, rootDir)
	log.Fatal(http.ListenAndServe(port, mux))
}
