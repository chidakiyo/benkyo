package lib

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"github.com/chidakiyo/benkyo/go-memleak-check/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func OfficialDatastore(g *gin.Context) {
	c := g.Request.Context()
	result := officialSearchDao(c)
	g.JSON(http.StatusOK, result)
}

func OfficialDatastorePointer(g *gin.Context){
	c := g.Request.Context()
	result := officialSearchDaoPointer(c)
	g.JSON(http.StatusOK, result)
}

func officialSearchDao(c context.Context) []Entity {
	var result []Entity
	ProjID := GetProject()
	client, err := datastore.NewClient(c, ProjID)
	if err != nil {
		log.Fatal(c, "client create error, %s", err.Error())
		//panic(err)
		return result
	}
	defer client.Close()

	q := datastore.NewQuery(KIND)
	_, err = client.GetAll(c, q, &result)
	if err != nil {
		log.Fatal(c, "fetch error %s", err.Error())
		return result
	}
	return result
}

func officialSearchDaoPointer(c context.Context) *[]Entity {
	var result []Entity
	ProjID := GetProject()
	client, err := datastore.NewClient(c, ProjID)
	if err != nil {
		log.Fatal(c, "client create error, %s", err.Error())
		//panic(err)
		return &result
	}
	defer client.Close()

	q := datastore.NewQuery(KIND)
	_, err = client.GetAll(c, q, &result)
	if err != nil {
		log.Fatal(c, "fetch error %s", err.Error())
		return &result
	}
	return &result
}

func DeleteAll(g *gin.Context) {
	c := g.Request.Context()
	ProjID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	client, err := datastore.NewClient(c, ProjID)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	q := datastore.NewQuery("kind1").KeysOnly()
	d := []interface{}{}
	keys, err := client.GetAll(c, q, d)
	if err != nil {
		fmt.Errorf("find error :%s\n", err.Error())
		g.String(http.StatusOK, "errorだよ")
		return
	}
	for _, k := range keys {
		fmt.Printf("target :%s\n", k.String())
		err := client.Delete(c, k)
		if err != nil {
			fmt.Errorf("delete fail :%s", err.Error())
		}
	}
	g.String(http.StatusOK, "finish")
}
