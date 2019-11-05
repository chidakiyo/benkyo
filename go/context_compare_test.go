package _go

import (
	"context"
	"github.com/gin-gonic/gin"
	"testing"
)

func Test_context_compare(t *testing.T) {

	pureContext := context.Background()
	ginContext := &gin.Context{}

	c(pureContext)
	c(ginContext)

	//g(pureContext)
	g(ginContext)

	context.WithCancel(ginContext)

}

func c(c context.Context){
}

func g(c *gin.Context) {
}

