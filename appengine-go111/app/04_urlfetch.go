package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
	"net/http"
)

func handleUrlFetch(g *gin.Context) {
	c := appengine.NewContext(g.Request)

	cl := urlfetch.Client(c)

	b := Benchmarker{}

	// 30回実行
	for i := 0; i <= 30 ; i++ {
		var err error
		b.Do(c, func() {
			_, err = cl.Get("https://chidakiyo.github.io?" + fmt.Sprintf("%d", i))
		})
		if err != nil {
			log.Errorf(c, "fetch error. %s", err)
			g.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	log.Infof(c, "fetch success.")
	log.Infof(c, "Result : %s", b.Result())

	// response
	g.String(http.StatusOK, "04_urlfetch")
}
