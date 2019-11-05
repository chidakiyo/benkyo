package lib

import (
	"context"
	"github.com/chidakiyo/benkyo/go-memleak-check/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// ContextLeakWithoutCancel context_withtimeoutを利用し、終了後にCancelを `実行しない` パターン
func ContextLeakWithoutCancel(g *gin.Context) {
	c := g.Request.Context()
	cc, _ := context.WithTimeout(g, 10*time.Minute)

	// ログを大量に出力してみる
	for i := 0; i < 100; i++ {
		log.Info(c, "log without: %v\n", cc)
	}

	g.String(http.StatusOK, "ok")
}

// ContextLeakWithCancel context_withtimeoutを利用し、終了後にCancelを実行するパターン
func ContextLeakWithCancel(g *gin.Context) {
	c := g.Request.Context()
	cc, cancel := context.WithTimeout(c, 10*time.Minute)
	defer cancel()

	// ログを大量に出力してみる
	for i := 0; i < 100; i++ {
		log.Info(c, "log with: %v\n", cc)
	}

	g.String(http.StatusOK, "ok")
}

// ContextLoop contextを1リクエスト中にループで大量に作成する
func ContextLoop(g *gin.Context) {
	c := g.Request.Context()
	var contextList []context.Context

	genChild(5000, &contextList, c)

	g.String(http.StatusOK, "ok")
}

// genChild 再帰的に子供のcontextを生成する
func genChild(counter int64, list *[]context.Context, parent context.Context) {
	child, _ := context.WithTimeout(parent, 10*time.Minute)
	*list = append(*list, child)
	log.Info(parent, "child generation: %d", counter)
	counter = counter - 1
	if counter <= 0 {
		return
	}
	genChild(counter, list, child)
}
