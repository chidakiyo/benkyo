package redis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"sync"
	"time"
)

var RedisPool *redis.Pool

func Redis_INCR(g *gin.Context) {
	base(g, func(c redis.Conn, i string) error {
		// インクリメント
		_, err := redis.Int(c.Do("INCR", "visits"))
		return err
	})
}

func Redis_EXISTS(g *gin.Context) {
	base(g, func(c redis.Conn, i string) error {
		// インクリメント
		_, err := redis.Int(c.Do("EXISTS", "visits"))
		return err
	})
}

func Redis_SET(g *gin.Context) {
	base(g, func(c redis.Conn, i string) error {
		// インクリメント
		_, err := redis.String(c.Do("SET", "k" + i, "value"))
		return err
	})
}

func Redis_GET(g *gin.Context) {
	base(g, func(c redis.Conn, i string) error {
		// インクリメント
		_, err := redis.String(c.Do("GET", "k" + i))
		return err
	})
}

func base(g *gin.Context, f func(c redis.Conn, i string) error) {
	//c := g.Request.Context()

	result := make(chan int64, 1000)
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		iString := fmt.Sprintf("%d", i)
		/*go*/ func() {
			defer wg.Done()
			conn := RedisPool.Get()
			defer conn.Close()
			start := time.Now().UnixNano()
			//_, err := redis.Int(conn.Do("INCR", "visits"))
			err := f(conn, iString)
			end := time.Now().UnixNano()

			t := (end-start)/int64(time.Microsecond)
			result <- t
			fmt.Printf("経過時間 %d microsec. \n", t)
			if err != nil {
				http.Error(g.Writer, "Error incrementing visitor counter", http.StatusInternalServerError)
				return
			}
		}()
	}
	wg.Wait()
	close(result)
	for w := range result {
		fmt.Fprintf(g.Writer, "%d\n", w)
	}
}
