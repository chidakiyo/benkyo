package main

import (
	"cloud.google.com/go/datastore"
	"cloud.google.com/go/profiler"
	"context"
	"fmt"
	"os"
)

type Entity struct {
	Name string `datastore:"Name"`
}

func main() {

	StartProfiler("leak-01", "get-entity")

	c := context.Background()

	ProjID := os.Getenv("PROJECT_ID")
	client, err := datastore.NewClient(c, ProjID)
	if err != nil {
		panic(err)
	}

	k := datastore.NameKey("kind1", "key1", nil)

	for {
		e := &Entity{}
		client.Get(c, k, e)

		fmt.Printf("Search at: %+v\n", e)
	}
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
