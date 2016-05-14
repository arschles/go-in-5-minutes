package models

import (
	"database/sql"
)

const (
	PersonTableName    = "person"
	PersonFirstNameCol = "first_name"
	PersonLastNameCol  = "last_name"
	PersonAgeCol       = "age"
)

// Person is the database model for a person
type Person struct {
	FirstName string
	LastName  string
	Age       uint
}

// CreatePersonTable uses db to create a new table for Person models, and returns the result
func CreatePersonTable(db *sql.DB) (sql.Result, error) {
	return db.Exec(
		"CREATE TABLE ? (? varchar(255), ? varchar(255), ? int)",
		PersonTableName,
		PersonFirstNameCol,
		PersonLastNameCol,
		PersonAgeCol,
	)
}
