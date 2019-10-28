package lib

import (
	"context"
	log2 "github.com/chidakiyo/benkyo/go-memleak-check/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ContextLeakWithoutCancel(g *gin.Context){
	c := g.Request.Context()
	log := log2.NewLogger(c)
	cc, _ := context.WithTimeout(c, 10 * time.Minute)
	for i:= 0;i < 100;i++ {
		log.Info("log without: %v\n", cc)
	}
	g.String(http.StatusOK, "ok")
}

func ContextLeakWithCancel(g *gin.Context) {
	c := g.Request.Context()
	log := log2.NewLogger(c)
	cc, cancel := context.WithTimeout(c, 10 * time.Minute)
	defer cancel()
	for i:= 0;i < 100;i++ {
		log.Info("log with: %v\n", cc)
	}
	g.String(http.StatusOK, "ok")
}

func ContextLoop(g *gin.Context) {
	c := g.Request.Context()
	//log := log2.NewLogger(c)
	var ccc []context.Context
	for i := 0; i < 10000; i++ {
		cc, _ := context.WithTimeout(c, 10 * time.Minute)
		ccc = append(ccc, cc)
		//log.Info("log loop: %v\n", cc)
	}
	g.String(http.StatusOK, "ok")
}
