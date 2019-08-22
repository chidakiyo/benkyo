package main

import (
	"cloud.google.com/go/profiler"
	"github.com/chidakiyo/benkyo/appengine-go111-profiler/handler"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"net/http"
)

func main() {
	if err := profiler.Start(profiler.Config{
		//Service:        projectId,
		//ServiceVersion: "1.0.0",
		// ProjectID must be set if not running on GCP.
		//ProjectID: "my-project",
		DebugLogging: true,
	}); err != nil {
		// TODO: Handle error.
	}

	route := gin.Default()
	http.Handle("/", route)
	handler.Handle(route)
	appengine.Main()
}
