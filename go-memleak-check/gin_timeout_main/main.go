package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	route := gin.Default()
	route.GET("long", func(g *gin.Context) {
		c, cancel := context.WithTimeout(g, 3*time.Second)
		defer cancel()

		ch := make(chan error, 1)

		go func() {
			time.Sleep(3 * time.Second)
			ch <- nil
		}()

		select {
		case <-c.Done():
			// timeout
			fmt.Println("<< timeout >>")
		case <-ch:
			fmt.Println("<< success >>")
		}

		g.String(http.StatusOK, "success")
	})
	http.Handle("/", route)
	route.Run(":8080")
}
