package note

import (
	"fmt"
	"log"
	"time"

	"github.com/Lachignol/martin-solving/database"
	erreur "github.com/Lachignol/martin-solving/error"
	modelofApp "github.com/Lachignol/martin-solving/model"
	_ "github.com/mattn/go-sqlite3"
)

func AddNote(newNoteTitle string) {
	stmt, _ := database.Db.Prepare("INSERT INTO digitalbrain (title,completed,created_at,completed_at) VALUES (?,?,?,?)")
	stmt.Exec(newNoteTitle, false, time.Now(), nil)
	defer stmt.Close()
	fmt.Printf("Nouvelle tache ajouté titre:%v \n", newNoteTitle)
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
			err := rows.Scan(&FindedNote.Id, &FindedNote.Title, &FindedNote.Completed, &FindedNote.Created_at, &FindedNote.Completed_at)
			if err != nil {
				log.Fatal(err)
			}

			if count == number {
				fmt.Println("Suppression en cours pour la note:", FindedNote.Id, FindedNote.Title, FindedNote.Completed, FindedNote.Created_at, FindedNote.Completed_at)
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
		err := rows.Scan(&FindedNote.Id, &FindedNote.Title, &FindedNote.Completed, &FindedNote.Created_at, &FindedNote.Completed_at)
		if err != nil {
			log.Fatal(err)
		}
		FindedNoteResult = append(FindedNoteResult, FindedNote)
	}
	err = rows.Err()
	erreur.CheckErr(err)
	return FindedNoteResult

}

func Edit(index int, newtitre string) error {
	var FindedNote modelofApp.Note
	var id int
	var numberofligneOfDb int

	err := database.Db.QueryRow("SELECT COUNT(*) FROM digitalbrain").Scan(&numberofligneOfDb)
	erreur.CheckErr(err)

	if numberofligneOfDb >= index {
		rows, err := database.Db.Query("select * from digitalbrain")
		erreur.CheckErr(err)
		defer rows.Close()
		count := 0
		for rows.Next() {
			count += 1
			err := rows.Scan(&FindedNote.Id, &FindedNote.Title, &FindedNote.Completed, &FindedNote.Created_at, &FindedNote.Completed_at)
			if err != nil {
				log.Fatal(err)
			}

			if count == index {
				id = FindedNote.Id
				rows.Close()
				break
			}

		}
		err = rows.Err()
		erreur.CheckErr(err)

		fmt.Print("Voulez-vous vraiment modifié cette note [y]/[n]? ")
		var response string
		_, err = fmt.Scanln(&response)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return err
		}
		fmt.Println("You entered:", response)
		if response == "yes" || response == "y" {
			stmt, err := database.Db.Prepare("UPDATE digitalbrain SET title = ? WHERE id = ?")
			if err != nil {
				fmt.Printf("Erreur lors de la préparation de la requête: %v\n", err)
				return err
			}
			defer stmt.Close()

			_, err = stmt.Exec(newtitre, id)
			if err != nil {
				fmt.Printf("Erreur lors de l'exécution de la requête: %v\n", err)
				return err
			}

			fmt.Printf("Le titre de la tâche avec l'ID %v a été mis à jour en: %v\n", id, newtitre)
		} else {
			log.Println("La note a été conservé ")
		}
	} else {
		log.Println("La note séléctioné n'existe pas ")
	}
	return nil
}

func Toggle(index int) error {
	var FindedNote modelofApp.Note
	var id int
	var numberofligneOfDb int
	var iscompleted bool
	var completed_at time.Time
	var completed bool

	err := database.Db.QueryRow("SELECT COUNT(*) FROM digitalbrain").Scan(&numberofligneOfDb)
	erreur.CheckErr(err)

	if numberofligneOfDb >= index {
		rows, err := database.Db.Query("select * from digitalbrain")
		erreur.CheckErr(err)
		defer rows.Close()
		count := 0
		for rows.Next() {
			count += 1
			err := rows.Scan(&FindedNote.Id, &FindedNote.Title, &FindedNote.Completed, &FindedNote.Created_at, &FindedNote.Completed_at)
			if err != nil {
				log.Fatal(err)
			}

			if count == index {
				id = FindedNote.Id
				iscompleted = FindedNote.Completed
				rows.Close()
				break
			}

		}
		err = rows.Err()
		erreur.CheckErr(err)

		stmt, err := database.Db.Prepare("UPDATE digitalbrain SET completed = ? ,completed_at = ? WHERE id = ?")
		if err != nil {
			fmt.Printf("Erreur lors de la préparation de la requête: %v\n", err)
			return err
		}
		defer stmt.Close()
		if !iscompleted {
			completed_at = time.Now()
			completed = !iscompleted
			_, err = stmt.Exec(completed, completed_at, id)
			if err != nil {
				fmt.Printf("Erreur lors de l'exécution de la requête: %v\n", err)
				return err
			}
		} else {
			completed = !iscompleted
			_, err = stmt.Exec(completed, nil, id)
			if err != nil {
				fmt.Printf("Erreur lors de l'exécution de la requête: %v\n", err)
				return err
			}
		}

		return nil
	}
	
	return nil
}
