package lib

import (
	"github.com/chidakiyo/benkyo/go-memleak-check/log"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func RandStream(g *gin.Context) {
	c := g.Request.Context()
	var buf []byte
	b := genBytes()
	buf = append(buf, *b...)
	log.Info(c, "byte: %v", buf)
	g.String(http.StatusOK, "%v", buf)
}

func genBytes() *[]byte {
	token := make([]byte, 8*8*8*8*8)
	rand.Read(token)
	return &token
}
