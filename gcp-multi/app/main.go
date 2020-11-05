package main

import (
	"github.com/chidakiyo/benkyo/gcp-multi/handler"
	"log"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.EnvHandler)

	port := "8080"
	if s := os.Getenv("PORT"); s != "" {
		port = s
	}

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("http.ListenAndServe: %v", err)
	}

}
