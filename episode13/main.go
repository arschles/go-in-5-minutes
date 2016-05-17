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
	// change this line to use a different database and connection string to connect to a different database
	db, err := sql.Open(sqlite3Str, memStr)
	if err != nil {
		log.Fatalf("error opening DB (%s)", err)
	}

	log.Printf("Creating new table")
	if _, crErr := models.CreatePersonTable(db); crErr != nil {
		log.Fatalf("Error creating table (%s)", crErr)
	}
	log.Printf("Created")

	me := models.Person{FirstName: "Aaron", LastName: "Schlesinger", Age: 29}
	log.Printf("Inserting %+v into the DB", me)
	if _, insErr := models.InsertPerson(db, me); insErr != nil {
		log.Fatalf("Error inserting new person into the DB (%s)", insErr)
	}
	log.Printf("Inserted")

	log.Printf("Selecting person from the DB")
	selectedMe := models.Person{}
	if err := models.SelectPerson(db, me.FirstName, me.LastName, me.Age, &selectedMe); err != nil {
		log.Fatalf("Error selecting person from the DB (%s)", err)
	}
	log.Printf("Selected %+v from the DB", selectedMe)

	log.Printf("Updating person in the DB")
	updatedMe := models.Person{
		FirstName: "Aaron",
		LastName:  "Schlesinger",
		Age:       30, // make this update after my birthday!
	}
	if err := models.UpdatePerson(db, selectedMe.FirstName, selectedMe.LastName, selectedMe.Age, updatedMe); err != nil {
		log.Fatalf("Error updating person in the DB (%s)", err)
	}

	log.Printf("Deleting person from the DB")
	if delErr := models.DeletePerson(db, selectedMe.FirstName, selectedMe.LastName, selectedMe.Age); delErr != nil {
		log.Fatalf("Error deleting person from the DB (%s)", delErr)
	}
	log.Printf("Deleted")

}
