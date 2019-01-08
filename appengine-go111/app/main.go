package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"net/http"
)

func main() {
	route := gin.Default()
	http.Handle("/", route)

	disableGinDebugLog()

	// ログの出力がリクエストスコープでまとまるか
	route.GET("/01", handleLog)

	appengine.Main() // Listen
}

func disableGinDebugLog(){
	// 本番に上げたときにはginのdebugログを出さない
	if !appengine.IsDevAppServer() {
		gin.SetMode(gin.ReleaseMode)
	}
}
