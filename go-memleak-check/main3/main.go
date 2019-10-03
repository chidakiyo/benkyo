package main

import (
	"cloud.google.com/go/datastore"
	"cloud.google.com/go/profiler"
	"context"
	"fmt"
	"os"
	"time"
)

type Entity struct {
	Name string `datastore:"Name"`
}

func main() {

	StartProfiler("leak-01", "create-client-and-get-03")

	c := context.Background()
	ProjID := os.Getenv("PROJECT_ID")

	for i := 0; i < 99; i++ {
		go func() {
			var counter int
			for {
				client, err := datastore.NewClient(c, ProjID)

				k := datastore.NameKey("kind1", "key1", nil)
				e := &Entity{}
				err = client.Get(c, k, e)
				if err != nil {
					fmt.Printf("error, %+v,  %d\n", err, counter)
					break
				}

				counter = counter + 1
				if err != nil {
					fmt.Printf("error, %+v,  %d\n", err, counter)
					break
				}
				fmt.Printf("search at: %+v\n", client)
			}
		}()
	}
	time.Sleep(1000 * time.Minute)
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
