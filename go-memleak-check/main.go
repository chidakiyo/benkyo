package main

import (
	"context"
	"github.com/rs/xid"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/clouddatastore"
	"os"
)

func main() {

	c := context.Background()

	datastore_(c)

}

func datastore_(c context.Context) {

	ProjID := os.Getenv("PROJECT_ID")
	client, err := clouddatastore.FromContext(c, datastore.WithProjectID(ProjID))
	if err != nil {
		panic(err)
	}

	xid := xid.New()

	// create key
	key := client.NameKey("kind1", xid.String(), nil)

	_, err = client.Put(c, key, struct {
		hoge string
		foo int
	}{
		hoge: "hogehoge",
		foo : 1,
	})
	if err != nil {
		panic(err)
	}

}
