package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"log"
	"net/http"
	"os"

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
		fmt.Printf("---------------------------------------------\n")
		fmt.Printf("%s\n", string(resp.Data))
		fmt.Printf("---------------------------------------------\n")
		s, _ := json.MarshalIndent(resp.Extensions, "", "  ")
		fmt.Printf("%s\n", s)
		fmt.Printf("---------------------------------------------\n")
		oc := graphql.GetOperationContext(ctx)
		fmt.Printf("%s\n", oc.RawQuery)


		return resp
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
