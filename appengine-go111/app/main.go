package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"math"
	"net/http"
	"time"
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

	// delay packageの動作確認
	route.GET("/06", handleTQDelay)

	appengine.Main() // Listen
}

func disableGinDebugLog(){
	// 本番に上げたときにはginのdebugログを出さない
	if !appengine.IsDevAppServer() {
		gin.SetMode(gin.ReleaseMode)
	}
}

type Benchmarker struct {
	results []int
}

func (b *Benchmarker) Do(c context.Context, f func()) {

	// ---- Do start ----
	start :=  time.Now()
	f()
	end := time.Now()
	// ---- Do end ----

	diff := end.Nanosecond() - start.Nanosecond()
	if b.results == nil {
		b.results = []int{}
	}
	b.results = append(b.results, diff)
	log.Infof(c, "time : %d msec. %d nanosec", diff / 1e6, diff)
}

func(b *Benchmarker) Result() string {
	max := 0
	min := math.MaxInt32
	var avg float64
	total := 0
	for _, v := range b.results {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
		total = total + v
	}
	avg = float64(total / 1e6) / float64(len(b.results))

	return fmt.Sprintf("max: %d, min: %d, avg: %f", max / 1e6, min / 1e6, avg)
}
