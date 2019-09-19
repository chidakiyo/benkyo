package main

import (
	"cloud.google.com/go/profiler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {

	startProfiler()
	route := gin.Default()
	http.Handle("/", route)

	route.GET("ds", mercariDatastore)
	route.GET("ods", officialDatastore)
	route.GET("ads", appengineDatastore)

	log.Fatal(http.ListenAndServe(":8080", nil))

	//appengine.Main()
}

func startProfiler() {
	if err := profiler.Start(profiler.Config{
		Service:        "leak-01",
		ServiceVersion: "0.0.2",
		// ProjectID must be set if not running on GCP.
		ProjectID:    os.Getenv("DATASTORE_PROJECT_ID"),
		DebugLogging: true,
	}); err != nil {
		panic(err)
	}
}
