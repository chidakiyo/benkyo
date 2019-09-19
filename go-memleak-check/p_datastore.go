package main

import (
	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
	"os"
)

func officialDatastore(g *gin.Context) {

	c := g.Request.Context()
	ProjID := os.Getenv("PROJECT_ID")
	client, err := datastore.NewClient(c, ProjID)
	if err != nil {
		panic(err)
	}

	xid := xid.New()

	// create key
	key := datastore.NameKey("kind1",xid.String(), nil )

	_, err = client.Put(c, key, &struct {
		hoge string
		foo  int
	}{
		hoge: "hogehoge",
		foo:  1,
	})
	if err != nil {
		panic(err)
	}

	g.String(http.StatusOK, "finish.")

}
