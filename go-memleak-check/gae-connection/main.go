package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()

	route.Run(":8080")
}



