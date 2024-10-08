package main

import (
	"context"
	"github.com/chidakiyo/benkyo/go-db-ent-atlas/ent"
	"log"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	diff()
	//migration()
}

func diff() {
	client, err := ent.Open("sqlite3", "file:ent.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	//
	// migration
	//
	ctx := context.Background()
	// Create a local migration directory.
	dir, err := migrate.NewLocalDir("migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Write migration diff.
	err = client.Schema.Diff(ctx, schema.WithDir(dir))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("complete.")
}

func migration() {
	client, err := ent.Open("sqlite3", "file:ent.db?_fk=1")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx,
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
