package lib

import (
	"context"
	"github.com/chidakiyo/benkyo/go-memleak-check/log"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/clouddatastore"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const KIND = "kind1"

// データを検索する
func MercariDatastoreSearchPointer(g *gin.Context) {
	c := g.Request.Context()
	result := searchDaoPointer(c)
	g.JSON(http.StatusOK, result)
}

func MercariDatastoreSearch(g *gin.Context){
	c := g.Request.Context()
	result := searchDao(c)
	g.JSON(http.StatusOK, result)
}

func searchDao(c context.Context) []Entity {
	var result []Entity
	ProjID := GetProject()
	client, err := clouddatastore.FromContext(c, datastore.WithProjectID(ProjID))
	if err != nil {
		log.Fatal(c, "connection error : %v", err)
		//panic(err)
		return result
	}
	q := client.NewQuery(KIND)
	client.GetAll(c, q, &result) // TODO error処理適当
	return result
}

func searchDaoPointer(c context.Context) *[]Entity {
	var result []Entity
	ProjID := GetProject()
	client, err := clouddatastore.FromContext(c, datastore.WithProjectID(ProjID))
	if err != nil {
		log.Fatal(c, "connection error : %v", err)
		//panic(err)
		return &result
	}
	q := client.NewQuery(KIND)
	client.GetAll(c, q, &result) // TODO error処理適当
	return &result
}

// データを投入する
func MercariDatastoreCreate(g *gin.Context) {
	c := g.Request.Context()
	ProjID := GetProject()
	client, err := clouddatastore.FromContext(c, datastore.WithProjectID(ProjID))
	if err != nil {
		log.Fatal(c, "connection error : %v", err)
		//panic(err)
		return
	}
	// ID生成
	xid := xid.New()
	// create key
	key := client.NameKey(KIND, xid.String(), nil)
	e := NewRandEntity()
	_, err = client.Put(c, key, &e)
	if err != nil {
		log.Fatal(c, "datastore put error : %v", err)
		//panic(err)
		return
	}
	g.String(http.StatusOK, "finish.")
}

func GetProject() string {
	if v := os.Getenv("GCP_PROJECT"); v != "" {
		return v
	} else if v := os.Getenv("GOOGLE_CLOUD_PROJECT"); v != "" {
		return v
	} else if v := os.Getenv("GCLOUD_PROJECT"); v != "" {
		return v
	}
	return ""
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomText(length int) string {
	const charSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(b)
}

func NewRandEntity() Entity {
	return Entity{
		Name: randomText(10),
		Foo:  rand.Int63(),
		Bar:  randomText(20),
	}
}

type Entity struct {
	Name string `datastore:"name"`
	Foo  int64  `datastore:"foo"`
	Bar  string `datastore:"bar"`
}
