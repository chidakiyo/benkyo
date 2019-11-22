package main

import (
	"context"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"flag"
	"fmt"
	redis_mod "github.com/chidakiyo/benkyo/appengine-memorystore/redis"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	r := gin.Default()

	redisHost := os.Getenv("REDISHOST")
	redisPort := os.Getenv("REDISPORT")

	redisServer := flag.String("redisServer", redisHost+":"+redisPort, "")
	redis_mod.RedisPool = newPool(*redisServer)

	// -- router --

	r.GET("/", func(g *gin.Context) {
		g.String(http.StatusOK, "ok")
	})

	r.GET("/r1", redis_mod.Redis_INCR)
	r.GET("/r2", redis_mod.Redis_EXISTS)
	r.GET("/r3", redis_mod.Redis_SET)
	r.GET("/r4", redis_mod.Redis_GET)

	// -- router --

	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: os.Getenv("GAE_PROJECT"),
	})
	if err != nil {
		fmt.Println("Stackdriver exporter initialize NG.")
		panic(err)
	}
	fmt.Println("Stackdriver exporter initialize OK.")
	trace.RegisterExporter(exporter)
	defer exporter.Flush()

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: &ochttp.Handler{
			Handler:     r,
			Propagation: &propagation.HTTPFormat{},
		},
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)
	<-sigCh
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
