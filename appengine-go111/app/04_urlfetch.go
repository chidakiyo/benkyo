package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
	"net/http"
)

func handleUrlFetch(g *gin.Context) {
	c := appengine.NewContext(g.Request)

	cl := urlfetch.Client(c)

	// urlfetch x3
	{
		_, err := cl.Get("https://chidakiyo.github.io")
		if err != nil {
			log.Errorf(c, "fetch error. %s", err)
			g.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	{
		_, err := cl.Get("https://chidakiyo.github.io")
		if err != nil {
			log.Errorf(c, "fetch error. %s", err)
			g.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	{
		_, err := cl.Get("https://chidakiyo.github.io")
		if err != nil {
			log.Errorf(c, "fetch error. %s", err)
			g.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	log.Infof(c, "fetch success.")

	// response
	g.String(http.StatusOK, "04_urlfetch")
}
