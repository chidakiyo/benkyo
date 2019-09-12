package main

import (
	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/profiler"
	"context"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"net/http"
	"os"
	"os/signal"
	"regexp"
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

	route.GET("/", middleware, handle)

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

func middleware(g *gin.Context) {
	// Fetch Authorization Header
	_, spanHeader := trace.StartSpan(g.Request.Context(), "header")
	bearerHeader := g.Request.Header.Get("Authorization")
	if bearerHeader == "" {
		g.AbortWithError(http.StatusUnauthorized, fmt.Errorf("No Authorization header found"))
		g.Abort()
		return
	}
	fmt.Printf("# BearerHeader: %s\n", bearerHeader)
	spanHeader.End()

	_, spanHeader2 := trace.StartSpan(g.Request.Context(), "header_parse")
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
	spanHeader2.End()
}

func verifyToken(c context.Context, g *gin.Context, bearerToken string) (IdTokenClaims, bool) {
	// Verify ID Token
	cc, spanVfy := trace.StartSpan(c, "verify_token")
	token, err := jwt.ParseWithClaims(bearerToken, &IdTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		kid := token.Header["kid"].(string)
		fmt.Printf("# kid: %s\n", kid)

		// Get certificate
		_, spanReq := trace.StartSpan(cc, "verify_request")
		resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		spanReq.End()

		_, spanDecode := trace.StartSpan(cc, "decode_certs")
		decoder := json.NewDecoder(resp.Body)
		var jsonBody interface{}
		if err := decoder.Decode(&jsonBody); err != nil {
			return nil, err
		}
		cert := jsonBody.(map[string]interface{})[kid].(string)
		spanDecode.End()

		fmt.Printf("# JsonBody: %+v\n", jsonBody)

		_, spanPemParse := trace.StartSpan(cc, "parse_pem")
		x,y := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		spanPemParse.End()

		return x, y
	})
	if err != nil {
		g.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid token: %s", err))
		return IdTokenClaims{}, false
	}
	fmt.Printf("# token: %+v\n", token)
	fmt.Printf("# token Claims: %+v\n", token.Claims)
	spanVfy.End()

	_, spanClaims := trace.StartSpan(c, "claim")
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
	spanClaims.End()

	return *claims, true
}

func handle(context *gin.Context) {

	bearerToken, _ := context.Get("TOKEN")
	ctx, spanHeader := trace.StartSpan(context.Request.Context(), "verify")
	claims, _ := verifyToken(ctx, context, bearerToken.(string))
	spanHeader.End()

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
