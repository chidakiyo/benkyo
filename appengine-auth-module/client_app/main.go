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

	route.GET("/", handle)
	appengine.Main()
}

func handle(context *gin.Context) {

	// ProjectID
	projectId := getProjectID()

	// audience
	audience := os.Getenv("ID_TOKEN_AUDIENCE")
	fmt.Printf("Audience: %s\n", audience)

	// ID_Token
	idToken := generateToken(audience)

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
}

func generateToken(audience string) string {
	idToken, err := metadata.Get("instance/service-accounts/default/identity?audience=" + audience)
	if err != nil {
		panic(err) // TODO 手抜き
	}
	fmt.Printf("ID Token: %s\n", idToken)
	return idToken
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
