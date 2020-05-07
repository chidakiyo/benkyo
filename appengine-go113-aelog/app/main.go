package main

import (
	"fmt"
	log "github.com/DeNA/aelog"
	"github.com/DeNA/aelog/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	route := gin.Default()

	route.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "[%s]", os.Getenv("MESSAGE"))
	})

	route.GET("/1", func(g *gin.Context) {
		ctx := g.Request.Context()
		log.Infof(ctx, "some log message")

		log.Debugf(ctx, "debug message")

		log.Errorf(ctx, "error message")

		log.Criticalf(ctx, "critical message")

		log.Warningf(ctx, "warning message")
	})

	h := middleware.AELogger("ServeHTTP")(route)
	p := os.Getenv("PORT")
	if len(p) == 0 {
		p = "8080"
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%s", p), h); err != nil {
		panic(err)
	}
}
