package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"math/rand"
	"os"
	"time"

	"github.com/chidakiyo/benkyo/gqlgen-todos2/graph/generated"
	"github.com/chidakiyo/benkyo/gqlgen-todos2/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {

	rMeta := RequestMetaFromContext(ctx)
	fmt.Printf("Meta : %+v\n", rMeta)

	oc := graphql.GetOperationContext(ctx)

	Output(&AuditLog{
		Cookie:    "some_cookie",
		Action:    oc.OperationName,
		Mark:      Mark,
		Trace:     fmt.Sprintf("projects/%s/traces/%s", os.Getenv("GCP_PROJECT"), rMeta.TraceID),
		ClientIP:  rMeta.RemoteIP,
		Timestamp: time.Now(),
	})

	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
