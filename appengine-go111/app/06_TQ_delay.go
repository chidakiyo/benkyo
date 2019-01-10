package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/delay"
	"google.golang.org/appengine/log"
	"net/http"
)

func handleTQDelay(g *gin.Context) {
	c := appengine.NewContext(g.Request)

	log.Infof(c, "before delay")

	// 非同期で実行
	err := revoker.Call(c)
	if err != nil {
		log.Errorf(c, "task create error. %s", err)
		g.String(http.StatusInternalServerError, err.Error())
		return
	}

	log.Infof(c, "after delay")

	// response
	g.String(http.StatusOK, "06_TQ_delay")
}

var revoker = delay.Func("delay", func(c context.Context) {
	log.Infof(c, "exec delay")
})
