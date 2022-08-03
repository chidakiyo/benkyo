package main

import (
	"context"
	"fmt"
	"github.com/chidakiyo/benkyo/go-db-ent/ent"
	"github.com/chidakiyo/benkyo/go-db-ent/ent/user"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	ctx := context.Background()

	datasource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "", "", "", "", "") // TODO

	client, err := ent.Open("postgres", datasource)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	//if err := client.Schema.Create(ctx); err != nil {
	//	log.Fatalf("failed creating schema resources: %v", err)
	//}

	//u, _ := CreateUser(ctx, client)
	//log.Printf("%v", u)

	u2, err := QueryUser(ctx, client)
	log.Printf("%v", u2)

	//u3, err := CreateCars(ctx, client)
	//log.Printf("%v", u3)
	log.Printf("%v", err)
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Create().SetAge(30).SetName("a8m").Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("field creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Query().Where(user.Name("a8m")).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("faild querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {

	tesla, err := client.Car.Create().SetModel("Tesla").SetRegisteredAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", tesla)

	ford, err := client.Car.Create().SetModel("Ford").SetRegisteredAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", tesla)

	a8m, err := client.User.Create().SetAge(30).SetName("a8m").AddCars(tesla, ford).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", a8m)
	return a8m, nil
}
