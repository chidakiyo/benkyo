package main

import (
	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/profiler"
	"context"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// profiler
	if err := profiler.Start(profiler.Config{
		//DebugLogging: true,
	}); err != nil {
		panic("プロファイラの起動に失敗 : " + err.Error())
	}

	// trace
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		//ProjectID: os.Getenv("GOOGLE_CLOUD_PROJECT"),
	})
	if err != nil {
		fmt.Println("Stackdriver exporter initialize NG.")
		panic(err)
	}
	fmt.Println("Stackdriver exporter initialize OK.")
	trace.RegisterExporter(exporter)
	defer exporter.Flush()
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()}) // 毎回取得

	route := gin.Default()
	http.Handle("/", route)

	route.GET("/", handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: &ochttp.Handler{
			Handler:     route,
			Propagation: &propagation.HTTPFormat{},
		},
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)
	<-sigCh
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
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
	req, err := http.NewRequestWithContext(context.Request.Context(), "GET", fmt.Sprintf("https://server-dot-%s.appspot.com", projectId), nil)
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
