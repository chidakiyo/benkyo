package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/clouddatastore"
	"net/http"
	"os"
)

func MercariDatastore(g *gin.Context) {

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