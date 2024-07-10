package note

import (
	"fmt"
	"log"

	"github.com/Lachignol/martin-solving/database"
	erreur "github.com/Lachignol/martin-solving/error"
	"github.com/Lachignol/martin-solving/model"
	_ "github.com/mattn/go-sqlite3"
)

func AddNote(newNote modelofApp.Note) {
	stmt, _ := database.Db.Prepare("INSERT INTO digitalbrain (id,name,description) VALUES (?, ?, ?)")
	stmt.Exec(nil, newNote.Name, newNote.Description)
	defer stmt.Close()
	fmt.Printf("Nouvelle note ajouté titre:%v description:%v \n", newNote.Name, newNote.Description)
}

func DeleteNote(number int) {
	var FindedNote modelofApp.Note
	var id int
	var numberofligneOfDb int

	err := database.Db.QueryRow("SELECT COUNT(*) FROM digitalbrain").Scan(&numberofligneOfDb)
	erreur.CheckErr(err)

	if numberofligneOfDb >= number {
		rows, err := database.Db.Query("select * from digitalbrain")
		erreur.CheckErr(err)
		defer rows.Close()
		count := 0
		for rows.Next() {
			count += 1
			err := rows.Scan(&FindedNote.Id, &FindedNote.Name, &FindedNote.Description)
			if err != nil {
				log.Fatal(err)
			}

			if count == number {
				fmt.Println("Suppression en cours pour la note:", FindedNote.Id, FindedNote.Name, FindedNote.Description)
				id = FindedNote.Id
				rows.Close()
				break
			}

		}
		err = rows.Err()
		erreur.CheckErr(err)

		fmt.Print("Voulez-vous vraiment supprimé cette note [y]/[n]? ")
		var response string
		_, err = fmt.Scanln(&response)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		fmt.Println("You entered:", response)
		if response == "yes" || response == "y" {
			supp, _ := database.Db.Prepare("DELETE FROM digitalbrain WHERE ID = ?")
			defer supp.Close()
			_, err := supp.Exec(id)
			erreur.CheckErr(err)

			log.Println("Note supprimé")
		} else {
			log.Println("La note a été conservé ")
		}
	} else {
		log.Println("La note séléctioné n'existe pas ")
	}
}

func RecupNotes() []modelofApp.Note {
	var FindedNote modelofApp.Note
	var FindedNoteResult []modelofApp.Note
	rows, err := database.Db.Query("select * from digitalbrain ")
	erreur.CheckErr(err)
	defer rows.Close()
	count := 0
	for rows.Next() {
		count += 1
		err := rows.Scan(&FindedNote.Id, &FindedNote.Name, &FindedNote.Description)
		if err != nil {
			log.Fatal(err)
		}
		FindedNoteResult = append(FindedNoteResult, FindedNote)
	}
	err = rows.Err()
	erreur.CheckErr(err)
	return FindedNoteResult

}
