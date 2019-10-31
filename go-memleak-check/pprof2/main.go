package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	pprof.Register(route)
	route.GET("/", func(c *gin.Context) {c.String(http.StatusOK, "hello hogehoge.")})
	http.Handle("/", route)
	route.Run(":8080")
}
