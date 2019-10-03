package lib

import (
	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
	"os"
)

//var client *datastore.Client
//var initialized bool

func OfficialDatastore(g *gin.Context) {

	c := g.Request.Context()
	ProjID := os.Getenv("PROJECT_ID")
	//if !initialized {
	//	_client, err := datastore.NewClient(c, ProjID)
	//	if err != nil {
	//		panic(err)
	//	}
	//	client = _client
	//	initialized = true
	//}
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
