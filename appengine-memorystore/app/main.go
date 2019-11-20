package main

import (
	"flag"
	redis_mod "github.com/chidakiyo/benkyo/appengine-memorystore/redis"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"os"
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

	r.GET("/hoge", redis_mod.Redis01)

	// -- router --

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
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
