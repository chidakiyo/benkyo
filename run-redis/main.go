package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"net/http"
	"os"
	"time"
)

func main() {
	p := os.Getenv("PORT")
	if len(p) == 0 {
		p = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		host := r.Host
		fmt.Fprintf(w, "(%s) (%s)", host, path)
	})

	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		host := r.URL.Query().Get("h")
		port := r.URL.Query().Get("p")

		// redis 接続
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: "", // no password set
			DB:       0,  // use default DB

			PoolSize:     10,
			MinIdleConns: 10,
			ReadTimeout:  1 * time.Second,
			MaxRetries:   3,
			DialTimeout:  2 * time.Second,

			OnConnect: func(conn *redis.Conn) error {
				fmt.Println("-> Redis Connect") // stdoutに出力
				return nil
			},
		})

		key :=fmt.Sprintf("KEY_HOGEHOGE_test_%d", time.Now().Second())

		client.Set(key, key, 10 * time.Minute) // 10分後

		fmt.Fprintf(w, "%v", client)
	})

	err := http.ListenAndServe(":" + p, nil)
	if err != nil {
		panic(err)
	}
}