package models

import (
	"context"

	pg "github.com/go-pg/pg/v10"
)

// DBConfigURL  config url of database
var DBConfigURL *pg.DB

// InitDb to initialise database
func InitDb() {
	DBConfigURL = pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "password",
		Database: "postgres",
	})

	//Check if the database is running
	ctx := context.Background()

	if err := DBConfigURL.Ping(ctx); err != nil {
		panic(err)
	}
}
