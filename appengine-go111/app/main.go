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

	// AE Datastoreでputする
	route.GET("/02", handleAEDatastore)

	// AE Datastoreでgetする
	route.GET("/02/get", handleAEDatastoreRead)

	// Cloud Datastoreでputするj
	route.GET("/03", handleCloudDatastore)

	// Cloud Datastoreでreadする
	route.GET("/03/get", handleCloudDatastoreRead)

	// urlfetchのtraceの確認
	route.GET("/04", handleUrlFetch)

	// http getのtraceの確認
	route.GET("/05", handleHttpGet)

	// delay packageの動作確認
	route.GET("/06", handleTQDelay)


	// ベンチマーク用
	route.GET("/91/init", handleInit)
	route.GET("/91", handlePerformance)

	// 大陸間ベンチマーク用
	route.GET("/92/run", handleEchoRun)
	route.GET("/92", handleEcho)

	appengine.Main() // Listen
}

func disableGinDebugLog(){
	// 本番に上げたときにはginのdebugログを出さない
	if !appengine.IsDevAppServer() {
		gin.SetMode(gin.ReleaseMode)
	}
}

type Benchmarker struct {
	results []int64
}

func (b *Benchmarker) Do(c context.Context, f func()) {

	// ---- Do start ----
	start :=  time.Now()
	f()
	diff := time.Since(start)
	// ---- Do end ----

	if b.results == nil {
		b.results = []int64{}
	}
	b.results = append(b.results, diff.Nanoseconds())
	log.Infof(c, "time : %d msec. %d nanosec", diff / 1e6, diff)
}

func(b *Benchmarker) Result() string {
	max := int64(0)
	min := int64(math.MaxInt64)
	var avg float64
	total := int64(0)
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
