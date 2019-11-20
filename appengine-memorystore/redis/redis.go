package redis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
)

var RedisPool *redis.Pool

func Redis01(g *gin.Context) {

	//c := g.Request.Context()

	conn := RedisPool.Get()
	defer conn.Close()

	counter, err := redis.Int(conn.Do("INCR", "visits"))
	if err != nil {
		http.Error(g.Writer, "Error incrementing visitor counter", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(g.Writer, "Visitor number: %d", counter)


	//conn, err := redis.Dial("tcp", "localhost:6379")
	//if err != nil {
	//	panic(err)
	//}
	//defer conn.Close()
	//
	//// 値の書き込み
	//r, err := conn.Do("SET", "temperature", "25")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(r) // OK
	//
	//// 値の読み出し
	//s, err := redis.String(conn.Do("GET", "temperature"))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(s) // 25

}