package models

import (
	"database/sql"
	"fmt"
)

const (
	// PersonTableName is the name of the table for the person model
	PersonTableName = "person"
	// PersonFirstNameCol is the column name of the model's first name
	PersonFirstNameCol = "first_name"
	// PersonLastNameCol is the column name of the model's last name
	PersonLastNameCol = "last_name"
	// PersonAgeCol is the column name of the model's age
	PersonAgeCol = "age"
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
		fmt.Sprintf("CREATE TABLE %s (%s varchar(255), %s varchar(255), %s int)",
			PersonTableName,
			PersonFirstNameCol,
			PersonLastNameCol,
			PersonAgeCol,
		),
	)
}

// InsertPerson inserts person into db
func InsertPerson(db *sql.DB, person Person) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("INSERT INTO %s VALUES(?, ?, ?)", PersonTableName),
		person.FirstName,
		person.LastName,
		person.Age,
	)
}

// SelectPerson selects a person with the given first & last names and age. On success, writes the result into result and on failure, returns a non-nil error and makes no modifications to result
func SelectPerson(db *sql.DB, firstName, lastName string, age uint, result *Person) error {
	row := db.QueryRow(
		fmt.Sprintf(
			"SELECT * FROM %s WHERE %s=? AND %s=? AND %s=?",
			PersonTableName,
			PersonFirstNameCol,
			PersonLastNameCol,
			PersonAgeCol,
		),
		firstName,
		lastName,
		age,
	)
	var retFirstName, retLastName string
	var retAge uint
	if err := row.Scan(&retFirstName, &retLastName, &retAge); err != nil {
		return err
	}
	result.FirstName = retFirstName
	result.LastName = retLastName
	result.Age = retAge
	return nil
}

// UpdatePerson updates the person with the given first & last names and age with newPerson. Returns a non-nil error if the update failed, and nil if the update succeeded
func UpdatePerson(db *sql.DB, firstName, lastName string, age uint, newPerson Person) error {
	_, err := db.Exec(
		fmt.Sprintf(
			"UPDATE %s SET %s=?,%s=?,%s=? WHERE %s=? AND %s=? AND %s=?",
			PersonTableName,
			PersonFirstNameCol,
			PersonLastNameCol,
			PersonAgeCol,
			PersonFirstNameCol,
			PersonLastNameCol,
			PersonAgeCol,
		),
		newPerson.FirstName,
		newPerson.LastName,
		newPerson.Age,
		firstName,
		lastName,
		age,
	)
	return err
}

// DeletePerson deletes the person with the given first & last names and age. Returns a non-nil error if the delete failed, and nil if the delete succeeded
func DeletePerson(db *sql.DB, firstName, lastName string, age uint) error {
	_, err := db.Exec(
		fmt.Sprintf(
			"DELETE FROM %s WHERE %s=? AND %s=? AND %s=?",
			PersonTableName,
			PersonFirstNameCol,
			PersonLastNameCol,
			PersonAgeCol,
		),
		firstName,
		lastName,
		age,
	)
	return err
}
