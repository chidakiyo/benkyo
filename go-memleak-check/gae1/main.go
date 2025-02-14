package main

import (
	"cloud.google.com/go/profiler"
	"context"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"fmt"
	"github.com/chidakiyo/benkyo/go-memleak-check/lib"
	"github.com/chidakiyo/benkyo/go-memleak-check/log"
	"github.com/gin-contrib/pprof" // see. https://qiita.com/junpayment/items/23934772f5385304eba8
	"github.com/gin-gonic/gin"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"google.golang.org/appengine"
	"net/http"
	//_ "net/http/pprof" // pprof
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	Bootstrap(func(e *gin.Engine) {
		// --------- Handler ----------
		// データストアにテスト用のデータを投入する
		e.GET("md/create", lib.MercariDatastoreCreate)
		e.GET("md/search_p", lib.MercariDatastoreSearchPointer)
		e.GET("md/search", lib.MercariDatastoreSearch)
		e.GET("md/search_nd", lib.MercariDatastoreSearchNoDeprecate)

		e.GET("od/search", lib.OfficialDatastore)
		e.GET("od/search_p", lib.OfficialDatastorePointer)

		e.GET("ads", lib.AppengineDatastore)
		e.GET("delete_all", lib.DeleteAll)

		e.GET("context_with", lib.ContextLeakWithCancel)
		e.GET("context_without", lib.ContextLeakWithoutCancel)
		e.GET("context_loop", lib.ContextLoop)

		e.GET("rand_01", lib.RandStream)
		e.GET("rand_02", lib.RandStreamString)

		e.GET("goro", func(i *gin.Context) {
			i.String(http.StatusOK, "goro :[%s] %d", os.Getenv("GAE_INSTANCE"), runtime.NumGoroutine())
		})
		e.GET("gome", func(i *gin.Context) {
			var mem runtime.MemStats
			runtime.ReadMemStats(&mem)
			i.String(http.StatusOK, "gome :[%s] \nalloc: %d, \ntotal_alloc: %d, \nheap_alloc: %d, \nheap_sys: %d\n", os.Getenv("GAE_INSTANCE"), mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
		})

		pprof.Register(e)
		// --------- Handler ----------
	},
		func() {}, // noop
		LoggingMiddleware(os.Getenv("GAE_SERVICE")),
	).Serve()
}

// Bootstrap はmicroservice(appengine以外でも)を起動する共通のmain処理
func Bootstrap(routing func(e *gin.Engine), conf func(), middleware ...gin.HandlerFunc) *Bootstrapper {

	if err := profiler.Start(profiler.Config{
		//DebugLogging: true,
	}); err != nil {
		//panic("プロファイラの起動に失敗 : " + err.Error())
		fmt.Printf("プロファイラは起動していません :%s", err.Error())
	}

	//gin.SetMode(gin.ReleaseMode) // debug出力を抑制
	r := gin.Default()
	r.Use(middleware...)
	http.Handle("/", r)

	conf() // noop
	fmt.Println("config initialize OK.")

	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID:GetProjectID(),
	})
	if err != nil {
		fmt.Println("Stackdriver exporter initialize NG.")
		//panic(err)
		fmt.Printf("exporterは起動していません :%s", err.Error())
	}
	fmt.Println("Stackdriver exporter initialize OK.")
	trace.RegisterExporter(exporter)
	//defer exporter.Flush()

	service := os.Getenv("GAE_SERVICE")
	r.GET("/", func(i *gin.Context) { i.String(http.StatusOK, service) }) // for debug

	// for /_ah/warmup
	r.GET("/_ah/warmup", func(i *gin.Context) {
		fmt.Printf("|| WARMUP REACHED ||\n")
		i.String(http.StatusOK, "warmup")
	})

	// --------- Handler ----------
	routing(r)
	// --------- Handler ----------

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: &ochttp.Handler{
			Handler:     r,
			Propagation: &propagation.HTTPFormat{},
		},
	}

	appengine.Main()

	return &Bootstrapper{
		server: server,
	}
}

type Bootstrapper struct {
	server *http.Server
}

func (b *Bootstrapper) Serve() {
	go func() {
		if err := b.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)
	<-sigCh
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := b.server.Shutdown(ctx); err != nil {
		panic(err)
	}
}

func LoggingMiddleware(module string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, logger, w := log.NewRequestLogger(c.Request.Context(), module, c.Request, c.Writer)
		defer logger.Finish()
		c.Request = c.Request.WithContext(ctx)
		//(c.Writer).(http.ResponseWriter) = w
		//c.Writer
		//c.Writer = gin.ResponseWriter((w).(gin.ResponseWriter))
		gin.DefaultWriter = w // ginの標準ログ

		c.Next()
	}
}

func GetProjectID() string {
	if v := os.Getenv("GCP_PROJECT"); v != "" {
		return v
	} else if v := os.Getenv("GOOGLE_CLOUD_PROJECT"); v != "" {
		return v
	} else if v := os.Getenv("GCLOUD_PROJECT"); v != "" {
		return v
	}
	return ""
}
