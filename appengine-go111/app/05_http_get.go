package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func handleHttpGet(g *gin.Context) {
	c := appengine.NewContext(g.Request)

	// http get x3
	{
		_, err := http.Get("https://chidakiyo.github.io")
		if err != nil {
			log.Errorf(c, "get error. %s", err)
			g.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	{
		_, err := http.Get("https://chidakiyo.github.io")
		if err != nil {
			log.Errorf(c, "get error. %s", err)
			g.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	{
		_, err := http.Get("https://chidakiyo.github.io")
		if err != nil {
			log.Errorf(c, "get error. %s", err)
			g.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	log.Infof(c, "http get success.")

	// response
	g.String(http.StatusOK, "05_http_get")
}
