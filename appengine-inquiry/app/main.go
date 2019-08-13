package main

import (
	"github.com/chidakiyo/benkyo/appengine-inquiry/module"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"net/http"
)

func main() {
	route := gin.Default()

	route.GET("/path", module.Path)

	http.Handle("/", route)
	appengine.Main()
}
