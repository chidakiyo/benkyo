package go_ratelimit

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/go-redis/redis_rate/v8"
	"os"
	"testing"
	"time"
)

func Test_テスト1(t *testing.T) {

	rh := os.Getenv("REDIS_HOST")

	rdb := redis.NewClient(&redis.Options{
		Addr: rh + ":6379",
	})
	_ = rdb.FlushDB().Err()

	limiter := redis_rate.NewLimiter(rdb)
	for i := 0; i < 100; i++ {
		fmt.Printf("counter %d\n", i)
		//DO:
		//for {
			res, err := limiter.Allow("project:123", redis_rate.PerSecond(10))
			if err != nil {
				panic(err)
			}
			fmt.Println("~~~~~~~", res.Allowed, res.Remaining, res.RetryAfter, res.ResetAfter)
			if res.Allowed == true {
				fmt.Println("				DO!")
				//break DO
			} else {
				time.Sleep(res.RetryAfter)
				fmt.Println("				DO!")
			}
			//time.Sleep(100 * time.Millisecond)
		//}
	}
}
