package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func handleLog(g *gin.Context) {
	c := appengine.NewContext(g.Request)

	log.Infof(c, "log info 1")
	log.Infof(c, "log info 2")

	log.Errorf(c, "log error 1")
	log.Errorf(c, "log error 2")

	log.Warningf(c, "log warning 1")
	log.Warningf(c, "log warning 2")

	log.Debugf(c, "log debug 1")
	log.Debugf(c, "log debug 2")

	log.Criticalf(c, "log critical 1")
	log.Criticalf(c, "log critical 2")

	log.Infof(c, "log info 3")

	// response
	g.String(http.StatusOK, "01_log")
}
