package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"os"
)

func main() {

	c := context.Background()

	ProjID := os.Getenv("PROJECT_ID")
	var counter int
	for {
		client, err := datastore.NewClient(c, ProjID)

		//k := datastore.NameKey("kind1", "key1", nil)
		//e := &Entity{}
		//client.Get(c, k, e)

		counter = counter + 1
		if err != nil {
			fmt.Printf("error, %+v,  %d\n", err, counter)
			break
		}
		fmt.Printf("deleted at: %+v\n", client)
	}
}
