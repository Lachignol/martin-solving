package database

import (
	"database/sql"
	_ "embed"
	"fmt"
	"os"

	"log"

	erreur "github.com/Lachignol/martin-solving/error"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

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
	Db, err = sql.Open("sqlite3", dbPath)
	erreur.CheckErr(err)
	return Db.Ping()

}

func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS digitalbrain (
	"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "title" TEXT NOT NULL,
    "completed" BOOLEAN NOT NULL,
    "created_at" DATETIME NOT NULL,
    "completed_at" DATETIME
	);`
	statement, err := Db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()

}

func Uninstall() {
	userdire, _ := os.UserHomeDir()
	dirpath := fmt.Sprintf("%s/databaseForMartinSolving/", userdire)
	dbPath := fmt.Sprintf("%s/databaseForMartinSolving/sqlite-digitalBrain.db", userdire)

	err := os.Remove(dbPath)
	if err != nil {
		log.Fatal(err)

	}

	log.Printf("Supression du fichier %v \n\n", dbPath)

	e := os.Remove(dirpath)
	if e != nil {
		log.Fatal(e)

	}
	log.Printf("Supression du repertoire %v\n\nIl ne vous reste plus qu'a supprimer le binaire 'martin-solving' dans le repertoire $HOME/go/bin/ .Pour y accéder,tapez 'cd $HOME/go/bin/'.\n", dirpath)

}
