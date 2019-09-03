package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"net/http"
)

func main() {
	route := gin.Default()
	http.Handle("/", route)

	route.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})

	appengine.Main()
}
