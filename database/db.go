package database

import (
	"database/sql"

	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDb() error {
	var err error
	db, err = sql.Open("sqlite3", "./sqlite-digitalBrain.db")
	checkErr(err)
	return db.Ping()
}

func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS digitalbrain (
	"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"name" TEXT,
	"description" TEXT
	);`
	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Table digitalbrain crée")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
