package types

import (
	pg "github.com/go-pg/pg/v10"
)

//Create create a new user and insert into db
func (u *User) Create(db *pg.DB) error {
	_, err := db.Model(u).Insert()
	return err
}

//Update update existing user
func (u *User) Update(db *pg.DB) error {
	_, err := db.Model(u).Update()
	return err
}

//Delete delete user
func (u *User) Delete(db *pg.DB) error {
	_, err := db.Model(u).Delete()
	return err
}
