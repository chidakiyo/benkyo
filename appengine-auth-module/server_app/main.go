package main

import (
	"cloud.google.com/go/compute/metadata"
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"net/http"
	"regexp"
)

func main() {
	route := gin.Default()
	http.Handle("/", route)

	route.GET("/", func(context *gin.Context) {
		// Fetch Authorization Header
		bearerHeader := context.Request.Header.Get("Authorization")
		if bearerHeader == "" {
			context.AbortWithError(http.StatusUnauthorized, fmt.Errorf("No Authorization header found"))
			return
		}

		re := regexp.MustCompile(`^\s*Bearer\s+(.+)$`)
		matched := re.FindStringSubmatch(bearerHeader)
		if len(matched) != 2 {
			context.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Authorization header is invalid format"))
			return
		}
		bearerToken := matched[1]

		// Verify ID Token
		token, err := jwt.ParseWithClaims(bearerToken, &IdTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			kid := token.Header["kid"].(string)

			// Get certificate
			resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			var jsonBody interface{}
			if err := decoder.Decode(&jsonBody); err != nil {
				return nil, err
			}
			cert := jsonBody.(map[string]interface{})[kid].(string)

			return jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		})
		if err != nil {
			context.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid token: %s", err))
			return
		}

		claims, ok := token.Claims.(*IdTokenClaims)
		if !(ok && token.Valid) {
			context.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid token"))
			return
		}

		// Get Project ID
		projectId, err := metadata.ProjectID()
		if err != nil {
			fmt.Println(err)
			context.String(http.StatusInternalServerError, "Error")
			return
		}

		// Check if the request came from same application
		if claims.Email != fmt.Sprintf("%s@appspot.gserviceaccount.com", projectId) {
			context.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid token: email is invalid, %s\n", claims.Email))
			return
		}

		fmt.Printf("%#v\n", token)
		context.String(http.StatusOK, "Request by: %s", claims.Email)
	})

	appengine.Main()
}

type IdTokenClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
