package main

import (
	"flag"
	"fmt"
	"os"
	"context"
	"github.com/go-pg/migrations/v8"
	pg "github.com/go-pg/pg/v10"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
`

func main() {
	flag.Usage = usage
	flag.Parse()

	db := pg.Connect(&pg.Options{
		Addr: "0.0.0.0:5432",
		User:     "postgres",
		Password: "password",
		Database: "postgres",
	})
	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Println(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
