package redis2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"net/http"
	"os"
	"testing"
	"time"
)

var RedisClient *redis.Client

func init() {
	host := os.Getenv("REDISHOST")
	port := os.Getenv("REDISPORT")
	RedisClient = NewClientWithParam(host, port)
	fmt.Println("go-redis init.")
}

func A(g *gin.Context) {
	r1 := put()
	r2 := get()
	fmt.Printf("%s\n%s\n", r1,r2)
	g.String(http.StatusOK, "%s\n%s\n", r1, r2)
}

func B(g *gin.Context) {
	start := time.Now().UnixNano()
	cmd := RedisClient.Set(key, value, 0)
	end := time.Now().UnixNano()
	result := end - start
	if cmd != nil && cmd.Err() != nil {
		g.String(http.StatusOK, "ng %d", result/1e6)
	}
	g.String(http.StatusOK, "ok %d", result/1e6)
	//fmt.Printf("%d \n", result/1e6)
}

const (
	key   = "go-redi-key"
	value = "go-redi-value"
)

func put() testing.BenchmarkResult {
	result := testing.Benchmark(func(b *testing.B) {
		b.ResetTimer()
		for i := 1; i <= b.N; i++ {
			cmd := RedisClient.Set(key, value, 0)
			if cmd != nil && cmd.Err() != nil {
				b.Fatal(cmd.Err())
			}
		}
	})
	return result
}

func get() testing.BenchmarkResult {
	result := testing.Benchmark(func(b *testing.B) {
		b.ResetTimer()
		for i := 1; i <= b.N; i++ {
			cmd := RedisClient.Get(key)
			if cmd != nil && cmd.Err() != nil {
				b.Fatal(cmd.Err())
			}
		}
	})
	return result
}

func NewClientWithParam(host, port string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB

		PoolSize:10,
		MinIdleConns:10,

		ReadTimeout: 1 * time.Second,

		MaxRetries: 3,
		DialTimeout: 2 * time.Second,

		OnConnect: func(conn *redis.Conn) error {
			fmt.Printf("-> Redis Connect %v\n", conn) // stdoutに出力
			fmt.Printf("%v \n", conn.ClientList())
			return nil
		},
	})
	return client
}