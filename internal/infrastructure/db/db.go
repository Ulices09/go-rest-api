package db

import (
	"context"
	"fmt"
	"go-rest-api/ent"
	"go-rest-api/internal/config"
	"net/url"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

func New(config config.Config) (client *ent.Client, err error) {
	drv, err := getMySqlDriver(config)

	if err != nil {
		return
	}

	client = ent.NewClient(ent.Driver(drv)) // logs: ent.Debug()
	err = client.Schema.Create(context.Background())

	if err != nil {
		return
	}

	return
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
