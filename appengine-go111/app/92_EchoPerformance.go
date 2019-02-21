package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
	"net/http"
)

// handleEchoRun はurlを引数として受け取り、別リージョンへのリクエストを投げ、パフォーマンスを測定する
func handleEchoRun(g *gin.Context) {
	c := appengine.NewContext(g.Request)

	cl := urlfetch.Client(c)

	b := Benchmarker{}

	q, ok := g.GetQuery("u")

	if !ok{
		log.Errorf(c, "query missing.")
		g.String(http.StatusBadRequest, "query missing.")
		return
	}
	log.Infof(c, "Query : %s", q)

	// 30回実行
	for i := 0; i <= 30 ; i++ {
		var err error
		b.Do(c, func() {
			_, err = cl.Get(q)
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
	g.String(http.StatusOK, "92_performance/run")
}

// handleEcho はレスポンスを返すだけ。
func handleEcho(g *gin.Context) {
	c := appengine.NewContext(g.Request)
	log.Infof(c, "response return")
	g.String(http.StatusOK, "ok!")
}


