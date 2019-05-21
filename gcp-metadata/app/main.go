package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"net/http"
	"os"
)

func main() {
	route := gin.Default()
	http.Handle("/", route)

	route.GET("/", root)
	route.GET("/header", headers)
	route.GET("/env", env)

	appengine.Main() // Listen
}

func root(g *gin.Context) {

	fmt.Println("ACCESS!!!")

	g.String(http.StatusOK, "ok")
}

func headers(g *gin.Context) {

	for k, v := range g.Request.Header {
		fmt.Printf("%s, %v\n", k, v)
	}

	g.String(http.StatusOK, "ok")
}

func env(g *gin.Context) {

	for _, v := range os.Environ() {
		fmt.Printf("%v\n", v)
	}

	g.String(http.StatusOK, "ok")
}
