package main

import (
	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func handleCloudDatastore(g *gin.Context) {
	c := appengine.NewContext(g.Request)

	dsClient, err := datastore.NewClient(c, "chida-test-012")
	if err != nil {
		log.Errorf(c, "datastore client create error. %s", err)
		g.String(http.StatusInternalServerError, err.Error())
		return
	}

	k := datastore.NameKey("Benkyo", "BK1", nil)
	e := Benkyo{
		Name:"Benkyo-name",
	}

	_, err = dsClient.Put(c, k, &e)
	if err != nil {
		log.Errorf(c, "datastore put error. %s", err)
		g.String(http.StatusInternalServerError, err.Error())
		return
	}

	log.Infof(c, "datastore put success. %v", k)

	// response
	g.String(http.StatusOK, "03_Cloud_Datastore")
}