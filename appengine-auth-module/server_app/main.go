package main

import (
	"cloud.google.com/go/compute/metadata"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"net/http"
	"os"
	"regexp"
)

func main() {
	route := gin.Default()
	http.Handle("/", route)

	route.GET("/", middleware, handle)
	appengine.Main()
}

func middleware(g *gin.Context) {
	// Fetch Authorization Header
	bearerHeader := g.Request.Header.Get("Authorization")
	if bearerHeader == "" {
		g.AbortWithError(http.StatusUnauthorized, fmt.Errorf("No Authorization header found"))
		g.Abort()
		return
	}
	fmt.Printf("# BearerHeader: %s\n", bearerHeader)

	re := regexp.MustCompile(`^\s*Bearer\s+(.+)$`)
	matched := re.FindStringSubmatch(bearerHeader)
	if len(matched) != 2 {
		g.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Authorization header is invalid format"))
		g.Abort()
		return
	}
	fmt.Printf("# Matched Result: %v\n", matched)
	bearerToken := matched[1]
	g.Set("TOKEN", bearerToken) // set token.
}

func verifyToken(g *gin.Context, bearerToken string) (IdTokenClaims, bool) {
	// Verify ID Token
	token, err := jwt.ParseWithClaims(bearerToken, &IdTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		kid := token.Header["kid"].(string)
		fmt.Printf("# kid: %s\n", kid)

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

		fmt.Printf("# JsonBody: %+v\n", jsonBody)

		return jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	})
	if err != nil {
		g.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid token: %s", err))
		return IdTokenClaims{}, false
	}
	fmt.Printf("# token: %+v\n", token)
	fmt.Printf("# token Claims: %+v\n", token.Claims)

	claims, ok := token.Claims.(*IdTokenClaims)
	if !(ok && token.Valid) {
		g.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid token"))
		return IdTokenClaims{}, false
	}

	projectID := getProjectID()
	// Check if the request came from same application
	if claims.Email != fmt.Sprintf("%s@appspot.gserviceaccount.com", projectID) {
		g.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid token: email is invalid, %s\n", claims.Email))
		return IdTokenClaims{}, false
	}

	return *claims, true
}

func handle(context *gin.Context) {

	bearerToken, _ := context.Get("TOKEN")
	claims, _ := verifyToken(context, bearerToken.(string))

	fmt.Printf("# Claims: %#v\n", claims)
	context.String(http.StatusOK, "Request by: %s", claims.Email)
}

type IdTokenClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// project id の取得
// metaサーバから取得するより環境変数(GAEの場合）で取得したほうがパフォーマンス良さそうなので
func getProjectID() string {
	envProjID, ok := os.LookupEnv("PROJECT_ID")
	if ok {
		fmt.Printf("Project ID: %s (env)\n", envProjID)
		return envProjID
	}
	projectId, err := metadata.ProjectID()
	if err != nil {
		panic(err) // TODO 手抜き
	}
	fmt.Printf("Project ID: %s (meta)\n", envProjID)
	return projectId
}
