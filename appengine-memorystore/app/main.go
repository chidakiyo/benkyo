package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/", func(g *gin.Context) {
		g.String(http.StatusOK, "ok")
	})

}