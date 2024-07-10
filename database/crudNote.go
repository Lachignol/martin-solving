package database

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddNote(newNote Note) {
	stmt, _ := db.Prepare("INSERT INTO digitalbrain (id,name,description) VALUES (?, ?, ?)")
	stmt.Exec(nil, newNote.Name, newNote.Description)
	defer stmt.Close()
	fmt.Printf("Nouvelle note ajouté titre:%v description:%v \n", newNote.Name, newNote.Description)
}

func DeleteNote(number int) {
	var FindedNote Note
	var id int
	var numberofligneOfDb int

	err := db.QueryRow("SELECT COUNT(*) FROM digitalbrain").Scan(&numberofligneOfDb)
	checkErr(err)

	if numberofligneOfDb >= number {
		rows, err := db.Query("select * from digitalbrain")
		checkErr(err)
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
		checkErr(err)

		fmt.Print("Voulez-vous vraiment supprimé cette note [y]/[n]? ")
		var response string
		_, err = fmt.Scanln(&response)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		fmt.Println("You entered:", response)
		if response == "yes" || response == "y" {
			supp, _ := db.Prepare("DELETE FROM digitalbrain WHERE ID = ?")
			defer supp.Close()
			_, err := supp.Exec(id)
			checkErr(err)

			log.Println("Note supprimé")
		} else {
			log.Println("La note a été conservé ")
		}
	} else {
		log.Println("La note séléctioné n'existe pas ")
	}
}

func RecupNotes() []Note {
	var FindedNote Note
	var FindedNoteResult []Note
	rows, err := db.Query("select * from digitalbrain ")
	checkErr(err)
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
	checkErr(err)
	fmt.Printf("%v", FindedNoteResult)
	return FindedNoteResult

}
