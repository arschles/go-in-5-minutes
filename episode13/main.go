package main

import (
	"database/sql"
	"log"

	"github.com/arschles/go-in-5-minutes/episode13/models"
	_ "github.com/mxk/go-sqlite/sqlite3"
)

const (
	sqlite3Str = "sqlite3"
	memStr     = ":memory:"
)

func main() {
	db, err := sql.Open(sqlite3Str, memStr)
	if err != nil {
		log.Fatalf("error opening DB (%s)", err)
	}

	log.Printf("Creating new table")
	if _, crErr := models.CreatePersonTable(db); crErr != nil {
		log.Fatalf("Error creating table (%s)", crErr)
	}
	log.Printf("Created")

	person := models.Person{FirstName: "Aaron", LastName: "Schlesinger", Age: 29}
	log.Printf("Inserting %+v into the DB", person)
	if _, insErr := models.InsertPerson(db, person); insErr != nil {
		log.Fatalf("Error inserting new person into the DB (%s)", insErr)
	}
	log.Printf("Inserted")

	log.Printf("Selecting person from the DB")
	retPerson := models.Person{}
	if err := models.SelectPerson(db, person.FirstName, person.LastName, person.Age, &retPerson); err != nil {
		log.Fatalf("Error selecting person from the DB (%s)", err)
	}
	log.Printf("Selected %+v from the DB", retPerson)

}
