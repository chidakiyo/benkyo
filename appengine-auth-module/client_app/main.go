package main

import (
	"cloud.google.com/go/compute/metadata"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	route := gin.Default()
	http.Handle("/", route)

	route.GET("/", func(context *gin.Context) {
		// Get Project ID
		projectId, err := metadata.ProjectID()
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// Get ID Token
		audience := os.Getenv("ID_TOKEN_AUDIENCE")
		fmt.Printf("Audience: %s\n", audience)

		idToken, err := metadata.Get("instance/service-accounts/default/identity?audience=" + audience)
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		fmt.Printf("ID Token: %s\n", idToken)

		// Call backend service
		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("https://server-dot-%s.appspot.com", projectId), nil)
		req.Header.Add("Authorization", "Bearer "+idToken)
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		context.String(http.StatusOK, "Response from backend:\n  %s", string(b))
	})

	appengine.Main()
}
