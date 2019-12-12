package main

import (
	"context"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"fmt"
	"github.com/chidakiyo/benkyo/appengine-memorystore2/redis1"
	"github.com/chidakiyo/benkyo/appengine-memorystore2/redis2"
	"github.com/gin-gonic/gin"
	"go.opencensus.io/plugin/ochttp"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	r := gin.Default()

	// -- router --
	r.GET("/", func(g *gin.Context) {g.String(http.StatusOK, "ok")})

	r.GET("/r1_a", redis1.A)

	r.GET("/r2_a", redis2.A)
	r.GET("/r2_b", redis2.B)

	// -- router --

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

