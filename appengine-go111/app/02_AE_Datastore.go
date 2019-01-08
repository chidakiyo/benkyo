package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"net/http"
)

type Benkyo struct {
	Name string `datastore:"name"`
}

func handleAEDatastore(g *gin.Context) {
	c := appengine.NewContext(g.Request)

	k := datastore.NewKey(c, "Benkyo", "BK1", 0, nil)

	_, err := datastore.Put(c, k, &Benkyo{
		Name: "Benkyo-name",
	})
	if err != nil {
		log.Errorf(c, "datastore put error. %s", err)
		g.String(http.StatusInternalServerError, err.Error())
		return
	}

	log.Infof(c, "datastore put success. %v", k)

	// response
	g.String(http.StatusOK, "02_AE_Datastore")
}

func handleAEDatastoreRead(g *gin.Context) {
	c := appengine.NewContext(g.Request)

	k := datastore.NewKey(c, "Benkyo", "BK1", 0, nil)

	// 3回Datastoreから取得する
	{
		bk := &Benkyo{}
		err := datastore.Get(c, k, bk)
		if err != nil {
			log.Errorf(c, "datastore get error. %s", err)
			g.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	{
		bk := &Benkyo{}
		err := datastore.Get(c, k, bk)
		if err != nil {
			log.Errorf(c, "datastore get error. %s", err)
			g.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	{
		bk := &Benkyo{}
		err := datastore.Get(c, k, bk)
		if err != nil {
			log.Errorf(c, "datastore get error. %s", err)
			g.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	log.Infof(c, "datastore get success. %v", k)

	// response
	g.String(http.StatusOK, "02_AE_Datastore/read")
}
