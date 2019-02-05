package main

import (
	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

const PROJECT = "chida-90204-jp"

func handleInit(g *gin.Context) {
	c := appengine.NewContext(g.Request)

	dsClient, err := datastore.NewClient(c, PROJECT)
	if err != nil {
		log.Errorf(c, "datastore client create error. %s", err)
		g.String(http.StatusInternalServerError, err.Error())
		return
	}

	b := Benchmarker{}

	// データ書き込み
	k := datastore.NameKey("Benkyo", "BK", nil)
	e := Benkyo{
		Name: "Benkyo-name",
	}
	// benchmark実行
	b.Do(c, func() {
		_, err = dsClient.Put(c, k, &e)
	})
	if err != nil {
		log.Errorf(c, "datastore put error. %s", err)
		g.String(http.StatusInternalServerError, err.Error())
		return
	}

	log.Infof(c, "datastore put success.")
	log.Infof(c, "Result : %s", b.Result())

	// response
	g.String(http.StatusOK, "91_init")
}

func handlePerformance(g *gin.Context) {
	c := appengine.NewContext(g.Request)

	dsClient, err := datastore.NewClient(c, PROJECT)
	if err != nil {
		log.Errorf(c, "datastore client create error. %s", err)
		g.String(http.StatusInternalServerError, err.Error())
		return
	}

	b := Benchmarker{}

	// 30回実行
	for i := 0; i <= 10 ; i++ {
		k := datastore.NameKey("Benkyo", "BK", nil)
		e := Benkyo{}
		// benchmark実行
		b.Do(c, func() {
			err = dsClient.Get(c, k, &e)
		})
		if err != nil {
			log.Errorf(c, "datastore get error. %s", err)
			g.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	log.Infof(c, "datastore get success.")
	log.Infof(c, "Result : %s", b.Result())

	// response
	g.String(http.StatusOK, "03_Cloud_Datastore")
}
