package lib

import (
	"github.com/chidakiyo/benkyo/go-memleak-check/log"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

func RandStream(g *gin.Context) {
	c := g.Request.Context()

	bit := uint64(0)
	bitSt := g.Query("b")
	if bitSt == "" {
		bit = 8
	} else {
		bit, _ = strconv.ParseUint(bitSt, 10, 64)
	}

	var buf []byte
	b := genBytes(bit)
	buf = append(buf, *b...)
	log.Info(c, "byte: %v", len(buf))
	g.String(http.StatusOK, "%v", len(buf)) // responseはデータのサイズ
}

func RandStreamString(g *gin.Context) {
	c := g.Request.Context()

	bit := uint64(0)
	bitSt := g.Query("b")
	if bitSt == "" {
		bit = 8
	} else {
		bit, _ = strconv.ParseUint(bitSt, 10, 64)
	}

	var buf []byte
	b := genBytes(bit)
	buf = append(buf, *b...)
	log.Info(c, "byte: %v", len(buf))
	g.String(http.StatusOK, "%v", buf) // responseはデータそのものをstring化
}

func genBytes(bit uint64) *[]byte {
	token := make([]byte, 1 << bit)
	rand.Read(token)
	return &token
}
