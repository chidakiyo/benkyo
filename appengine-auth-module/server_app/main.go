package main

import (
	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/profiler"
	"context"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"crypto/rsa"
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
	"sync"
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

type TokenCache struct {
	PubKey rsa.PublicKey
	TTL    time.Time
}
var tokenCache = sync.Map{}

func verifyToken(c context.Context, g *gin.Context, bearerToken string) (IdTokenClaims, bool) {
	// Verify ID Token
	cc, spanVfy := trace.StartSpan(c, "verify_token")
	token, err := jwt.ParseWithClaims(bearerToken, &IdTokenClaims{}, func(parsedToken *jwt.Token) (interface{}, error) {
		kid := parsedToken.Header["kid"].(string)
		fmt.Printf("# ParsedToken: %+v\n", parsedToken)
		fmt.Printf("# ParsedClaims: %+v\n", parsedToken.Claims)
		fmt.Printf("# ParsedSigningMethod: %+v\n", parsedToken.Method)
		fmt.Printf("# kid: %s\n", kid)

		audience := parsedToken.Claims.(*IdTokenClaims).Audience
		_cachedToken, ok := tokenCache.Load(audience)
		cachedToken := _cachedToken.(TokenCache)
		if ok {
			key := cachedToken.PubKey
			return &key, nil
		} else {
			// Get certificate
			key, err := getCert(cc, kid)

			// cache put
			tokenCache.Store(audience,
				TokenCache{
					PubKey: *key,
					TTL:    time.Now(),
				})
			return key, err
		}
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

func getCert(ctx context.Context, kid string) (*rsa.PublicKey, error) {

	// Get certificate
	_, spanReq := trace.StartSpan(ctx, "verify_request")
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	spanReq.End()

	_, spanDecode := trace.StartSpan(ctx, "decode_certs")
	decoder := json.NewDecoder(resp.Body)
	var jsonBody interface{}
	if err := decoder.Decode(&jsonBody); err != nil {
		return nil, err
	}
	cert := jsonBody.(map[string]interface{})[kid].(string)
	spanDecode.End()

	fmt.Printf("# JsonBody: %+v\n", jsonBody)
	fmt.Printf("# Cert: %+v\n", cert)

	_, spanPemParse := trace.StartSpan(ctx, "parse_pem")
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	spanPemParse.End()

	return pubKey, err
}

func handle(context *gin.Context) {

	_bearerToken, _ := context.Get("TOKEN")
	bearerToken := _bearerToken.(string)
	fmt.Printf("Baarer Token: %s\n", bearerToken)

	//fmt.Printf("tokenCache size : %d\n", len(tokenCache))
	//if len(tokenCache) > 0 {
	//	for k, v := range tokenCache {
	//		fmt.Printf("==> key: %v\n", k)
	//		fmt.Printf("==> value: %v\n", v)
	//	}
	//}

	ctx, spanHeader := trace.StartSpan(context.Request.Context(), "verify")
	claims, _ := verifyToken(ctx, context, bearerToken)
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
