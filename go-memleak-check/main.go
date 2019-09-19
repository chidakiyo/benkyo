package main

import (
	"cloud.google.com/go/profiler"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/clouddatastore"
	"log"
	"net/http"
	"os"
)

func main() {

	startProfiler()
	route := gin.Default()
	http.Handle("/", route)

	route.GET("ds", datastoreHandle)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func startProfiler() {
	if err := profiler.Start(profiler.Config{
		Service:        "leak-01",
		ServiceVersion: "0.0.0",
		// ProjectID must be set if not running on GCP.
		ProjectID:    os.Getenv("DATASTORE_PROJECT_ID"),
		DebugLogging: true,
	}); err != nil {
		panic(err)
	}
}

func datastoreHandle(g *gin.Context) {

	c := g.Request.Context()
	ProjID := os.Getenv("PROJECT_ID")
	client, err := clouddatastore.FromContext(c, datastore.WithProjectID(ProjID))
	if err != nil {
		panic(err)
	}

	xid := xid.New()

	// create key
	key := client.NameKey("kind1", xid.String(), nil)

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
