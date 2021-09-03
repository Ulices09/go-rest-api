package db

import (
	"context"
	"go-rest-api/ent"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb() *ent.Client {
	db, err := ent.Open("sqlite3", "file:app.db?_fk=1")

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return db
}
