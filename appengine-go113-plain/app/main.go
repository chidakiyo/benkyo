package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	route := gin.Default()

	route.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "[%s]", os.Getenv("MESSAGE"))
	})

	p := os.Getenv("PORT")
	if len(p) == 0 {
		p = "8080"
	}
	route.Run(":" + p)
}