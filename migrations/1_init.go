package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	fmt.Println("init function")
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating tables...")
		_, err := db.Exec(`
		CREATE TABLE organization (
			id serial PRIMARY KEY,
			org_name varchar(30) NOT NULL,
			org_image bytea
		);
		CREATE TABLE users (
			id serial PRIMARY KEY,
			first_name varchar(30),
			last_name varchar(30),
			email varchar(50) UNIQUE NOT NULL,
			password varchar NOT NULL,
			role integer DEFAULT 0,
			createdAt timestamp
		);
		CREATE TABLE subscriptions (
			uid integer references organization(id),
			oid integer references users(id),
			CONSTRAINT PK PRIMARY KEY(uid, oid)
		);
		CREATE TABLE polls (
			id serial PRIMARY KEY,
			questions varchar(30) NOT NULL,
			options varchar(30)[] NOT NULL,
			opensat timestamp,
			closesat timestamp,
			uid integer references organization(id),
			createdAt timestamp
		);
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping tables...")
		_, err := db.Exec(`
		DROP TABLE subscriptions;
		DROP TABLE users;
		DROP TABLE organization;
		DROP TABLE polls;
		`)
		return err
	})
}
