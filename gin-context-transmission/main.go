package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	c := context.Background()
	c = context.WithValue(c, "test", "testtest")
	do(c)
}

func do(ctx context.Context) {
	g := gin.New()

	// 普通に
	g.GET("/1", func(g *gin.Context) {
		c := g.Request.Context()
		fmt.Printf("%+v", c)
		g.String(http.StatusOK, http.StatusText(http.StatusOK))
	})

	fff := func(g *gin.Context) {
		c := g.Request.Context()
		c = context.WithValue(c, "middle", "middlemiddle")
		g.Request = g.Request.WithContext(c)
	}

	// middlewareする
	g.GET("/2", fff, func(g *gin.Context) {
		c := g.Request.Context()
		fmt.Printf("%+v", c)
		g.String(http.StatusOK, http.StatusText(http.StatusOK))
	})

	g.Run(":8080")
}