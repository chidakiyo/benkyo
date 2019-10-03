package main

import (
	"cloud.google.com/go/profiler"
	"github.com/chidakiyo/benkyo/go-memleak-check/lib"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {

	StartProfiler("leak-01", "0.0.2")
	route := gin.Default()
	http.Handle("/", route)

	route.GET("ds", lib.MercariDatastore)
	route.GET("ods", lib.OfficialDatastore)
	route.GET("ads", lib.AppengineDatastore)

	log.Fatal(http.ListenAndServe(":8080", nil))

	//appengine.Main()
}

func StartProfiler(service, version string) {
	if err := profiler.Start(profiler.Config{
		Service:        service,
		ServiceVersion: version,
		// ProjectID must be set if not running on GCP.
		ProjectID:    os.Getenv("PROJECT_ID"),
		DebugLogging: true,
	}); err != nil {
		panic(err)
	}
}
