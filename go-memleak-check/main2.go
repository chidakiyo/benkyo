package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"os"
)

type Entity struct {
	Name string `datastore:"Name"`
}

func main() {

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

		fmt.Printf("deleted at: %+v\n", e)
	}
}
