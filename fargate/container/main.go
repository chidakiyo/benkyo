package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	verstion := os.Getenv("VERSION")
	fmt.Fprintf(w, "<h1>%s</h1>", verstion)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}