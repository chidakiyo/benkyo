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

	// AE Datastoreでputする際のtrace
	route.GET("/02", handleAEDatastore)

	// AE Datastoreでgetする際のtrace
	route.GET("/02/get", handleAEDatastoreRead)

	// Cloud Datastoreでputする際のtrace
	route.GET("/03", handleCloudDatastore)

	// urlfetchのtraceの確認
	route.GET("/04", handleUrlFetch)

	// http getのtraceの確認
	route.GET("/05", handleHttpGet)

	appengine.Main() // Listen
}

func disableGinDebugLog(){
	// 本番に上げたときにはginのdebugログを出さない
	if !appengine.IsDevAppServer() {
		gin.SetMode(gin.ReleaseMode)
	}
}
