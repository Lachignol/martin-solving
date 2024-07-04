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
	dirpath := fmt.Sprintf("%s/databaseForMartinSolving/", userdire)
	dbPath := fmt.Sprintf("%s/databaseForMartinSolving/sqlite-digitalBrain.db", userdire)

	// Vérifier si le fichier existe
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Printf("Le fichier de la base de données %s n'existe pas.", dbPath)
		log.Printf("Nous procédons donc a l'installation ...")
		//si il existe pas on crée le repertoire
		err := os.Mkdir(dirpath, 0700)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf("Création du repertoire %s .", dirpath)
		log.Printf("Installation terminée.")
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

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
