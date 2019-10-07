package main

import (
	"github.com/chidakiyo/benkyo/appengine-mail-receive/mail"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"net/http"
)

func main() {
	route := gin.Default()
	http.Handle("/", route)

	route.POST("/_ah/mail/:address", mail.Mail)
	appengine.Main()
}

