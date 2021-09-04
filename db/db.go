package db

import (
	"context"
	"go-rest-api/config"
	"go-rest-api/ent"
	"log"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDb(config config.Config) *ent.Client {
	drv, err := sql.Open("mysql", config.DB.Url)

	if err != nil {
		log.Fatalf("Failed opening connection to db: %v", err)
	}

	client := ent.NewClient(ent.Driver(drv))

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
