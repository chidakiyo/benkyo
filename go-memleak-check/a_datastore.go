package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

func appengineDatastore(g *gin.Context) {

	c := appengine.NewContext(g.Request)

	xid := xid.New()

	// create key
	key := datastore.NewKey(c, "kind1", xid.String(), 0, nil)

	_, err := datastore.Put(c, key, &struct {
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
