package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"net/http"
)

type Benkyo struct {
	Name string `datastore:"name"`
}

// 書き込みテスト
func handleAEDatastore(g *gin.Context) {
	c := appengine.NewContext(g.Request)

	b := Benchmarker{}

	// 30回実行
	for i:= 0; i< 30;i++{
		count := fmt.Sprintf("%d", i + 1)
		k := datastore.NewKey(c, "Benkyo", "BK_" + count, 0, nil)
		var err error
		// benchmark実行
		b.Do(c, func() {
			_, err = datastore.Put(c, k, &Benkyo{
				Name: "Benkyo-name",
			})
		})
		if err != nil {
			log.Errorf(c, "datastore put error. %s", err)
			g.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	log.Infof(c, "datastore put success.")
	log.Infof(c, "Result : %s", b.Result())

	// response
	g.String(http.StatusOK, "02_AE_Datastore")
}

// 読み出しテスト
func handleAEDatastoreRead(g *gin.Context) {
	c := appengine.NewContext(g.Request)

	k := datastore.NewKey(c, "Benkyo", "BK1", 0, nil)

	b := Benchmarker{}

	// 30回実行
	for i := 0; i <= 30 ; i++ {
		bk := &Benkyo{}
		var err error
		// benchmark実行
		b.Do(c, func() {
			err = datastore.Get(c, k, bk)
		})
		if err != nil {
			log.Errorf(c, "datastore get error. %s", err)
			g.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	log.Infof(c, "datastore get success. %v", k)
	log.Infof(c, "Result : %s", b.Result())

	// response
	g.String(http.StatusOK, "02_AE_Datastore/read")
}
