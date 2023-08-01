package main

import (
	"cloud.google.com/go/compute/metadata"
	"fmt"
	log "github.com/chidakiyo/gcplog"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	route()
	fmt.Println("Start")

	p, err := metadata.ProjectID()
	if err != nil {
		panic(err)
	}
	log.Initialize("test", p)
	panic(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func route() {
	http.HandleFunc("/log1", log1)
}

func log1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
	log.Structure.Infof(r.Context(), struct {
		Typ     string `json:"@type"`
		Message string `json:"message"`
	}{
		Typ:     "type.googleapis.com/google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent",
		Message: fmt.Sprintf("%s : %s", "アラートです", time.Now().GoString()),
	})
	w.WriteHeader(http.StatusOK)
}
