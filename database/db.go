package database

import (
	"database/sql"
	_ "embed"
	"fmt"
	"os"

	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDb() error {
	userdire, _ := os.UserHomeDir()
	dbPath := fmt.Sprintf("%s/go/src/github.com/Lachignol/martin-solving/sqlite-digitalBrain.db", userdire)

	// Vérifier si le fichier existe
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Fatalf("Le fichier de la base de données %s n'existe pas.", dbPath)
		return err
	}

	var err error
	db, err = sql.Open("sqlite3", dbPath)
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
