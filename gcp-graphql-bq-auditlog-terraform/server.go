package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/DeNA/aelog/middleware"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/chidakiyo/benkyo/gqlgen-todos2/graph"
	"github.com/chidakiyo/benkyo/gqlgen-todos2/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
		//a, _ := json.MarshalIndent(ctx, "", "  ")
		//fmt.Printf("---------------------------------------------\n")
		//fmt.Printf("%+v\n", ctx)
		//fmt.Printf("---------------------------------------------\n")
		return next(ctx)
	})

	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		//fmt.Printf("---------------------------------------------\n")
		//fmt.Printf("%+v\n", ctx)
		//fmt.Printf("---------------------------------------------\n")
		resp := next(ctx)
		return resp
	})

	srv.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		resp := next(ctx)
		//fmt.Printf("---------------------------------------------\n")
		//fmt.Printf("%s\n", string(resp.Data))
		//fmt.Printf("---------------------------------------------\n")
		//s, _ := json.MarshalIndent(resp.Extensions, "", "  ")
		//fmt.Printf("%s\n", s)
		//fmt.Printf("---------------------------------------------\n")
		//oc := graphql.GetOperationContext(ctx)
		//fmt.Printf("%s\n", oc.RawQuery)
		return resp
	})

	router := chi.NewRouter()
	router.Use(middleware.AELogger("test"))
	router.Use(graph.RequestMiddleware())

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	sv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	go func() {
		if err := sv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	log.Printf("SIG: %d \n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := sv.Shutdown(ctx); err != nil {
		panic("err")
	}
}
