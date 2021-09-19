package db

import (
	"context"
	"fmt"
	"go-rest-api/ent"
	"go-rest-api/internal/config"
	"log"
	"net/url"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDb(config config.Config) *ent.Client {
	drv, err := getMySqlDriver(config)

	if err != nil {
		log.Fatalf("Failed opening connection to db: %v", err)
	}

	client := ent.NewClient(ent.Driver(drv)) // logs: ent.Debug()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func getMySqlDriver(c config.Config) (*sql.Driver, error) {
	connection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		c.DB.User,
		c.DB.Password,
		c.DB.Host,
		c.DB.Port,
		c.DB.Name,
	)

	mysqlParams := url.Values{}
	mysqlParams.Add("parseTime", "true")
	mysqlParams.Add("loc", "America/Lima")

	dsn := fmt.Sprintf("%s?%s", connection, mysqlParams.Encode())

	return sql.Open("mysql", dsn)
}
