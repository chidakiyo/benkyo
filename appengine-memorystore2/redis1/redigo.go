package redis1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"os"
	"testing"
	"time"
)

var RedisPool *redis.Pool

func init() {
	host := os.Getenv("REDISHOST")
	port := os.Getenv("REDISPORT")
	RedisPool = newPool(fmt.Sprintf("%s:%s", host, port))
	fmt.Println("redigo init.")
}

func A(g *gin.Context) {
	r1 := put()
	r2 := get()
	fmt.Printf("%s\n%s\n", r1,r2)
	g.String(http.StatusOK, "%s\n%s\n", r1, r2)
}

const (
	key   = "redigo-key"
	value = "redigo-value"
)

func put() testing.BenchmarkResult {

	conn := RedisPool.Get()
	defer conn.Close()

	result := testing.Benchmark(func(b *testing.B) {
		b.ResetTimer()
		for i := 1; i <= b.N; i++ {
			_, err := conn.Do("SET", key, value)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	return result
}

func get() testing.BenchmarkResult {
	conn := RedisPool.Get()
	defer conn.Close()

	result := testing.Benchmark(func(b *testing.B) {
		b.ResetTimer()
		for i := 1; i <= b.N; i++ {
			_, err := redis.String(conn.Do("GET", key))
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	return result
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
