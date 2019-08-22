package handler

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func Handle(e *gin.Engine) {
	// base
	e.GET("/", func(g *gin.Context) {
		g.String(http.StatusOK, "ok")
	})

	// loop
	e.GET("/2", func(g *gin.Context) {
		c := appengine.NewContext(g.Request)
		for i := 0; i <= 10000; i++ {
			log.Infof(c, "num: %d", i)
		}
		g.String(http.StatusOK, "ok")
	})
}
