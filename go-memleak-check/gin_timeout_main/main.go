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

	route.GET("01", func(g *gin.Context) {
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

	route.GET("02", func(i *gin.Context) {

		ch := make(chan error, 1)

		// main process
		go func() {
			time.Sleep(4 * time.Second)
			ch <- nil // complete
		}()

		select {
		case <-ch: // success
			i.String(http.StatusOK, "complete") // success
			return
		case <-time.After(2 * time.Second): // timeout
			i.String(http.StatusGatewayTimeout, "ごめんね")
			return
		}

	})

	http.Handle("/", route)
	route.Run(":8080")
}
