package lib

import (
	"cloud.google.com/go/datastore"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
	"os"
)

func OfficialDatastore(g *gin.Context) {

	c := g.Request.Context()
	ProjID := os.Getenv("PROJECT_ID")
	client, err := datastore.NewClient(c, ProjID)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	xid := xid.New()

	// create key
	key := datastore.NameKey("kind1", xid.String(), nil)

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
