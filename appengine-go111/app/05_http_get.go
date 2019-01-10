package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func handleHttpGet(g *gin.Context) {
	c := appengine.NewContext(g.Request)

	b := Benchmarker{}

	// 30回実行
	for i := 0; i <= 30 ; i++ {

		var err error
		b.Do(c, func() {
			_, err = http.Get("https://chidakiyo.github.io?" + fmt.Sprintf("%d", i))
		})
		if err != nil {
			log.Errorf(c, "get error. %s", err)
			g.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	log.Infof(c, "http get success.")
	log.Infof(c, "Result : %s", b.Result())

	// response
	g.String(http.StatusOK, "05_http_get")
}
