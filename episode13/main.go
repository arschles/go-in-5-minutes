package main

import (
	"database/sql"
	_ "github.com/mxk/go-sqlite/sqlite3"
	"log"
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

}
