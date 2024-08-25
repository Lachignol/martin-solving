/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"errors"
	"log"

	"strings"
	"time"

	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/spf13/cobra"
)


var inputValue string
var outputValue string
var replaceValue string
var replaceByValue string

// actionPageCmd represents the actionPage command
var actionOnFileCmd = &cobra.Command{
	Use:   "actionOnFile",
	Short: "Permet de remplacer un pattern dans fichier en entrée et crée un fichier en sortie",
	Long:  "Cette commande permet de remplacer un motif spécifique dans un fichier d'entrée et de créer un nouveau fichier en sortie contenant le contenu modifié. Elle est particulièrement utile pour les utilisateurs qui souhaitent effectuer des remplacements de texte dans des fichiers sans modifier l'original.",
	Args:  cobra.MaximumNArgs(4),

	Run: func(cmd *cobra.Command, args []string) {

		 if inputValue == "" && outputValue == "" && replaceValue == "" && replaceByValue == "" {
			fp := filepicker.New()
			fp.DirAllowed = false
			fp.AllowedTypes = []string{".txt"}
			// pour definire le repertoire courant au repertoire user  remplacer par :
			// fp.CurrentDirectory = "./"
			fp.CurrentDirectory, _ = os.UserHomeDir()

			m := modelFilepicker{
				filepicker: fp,
			}
			tm, _ := tea.NewProgram(&m).Run()
			mm := tm.(modelFilepicker)

			file := mm.selectedFile

			if file != "" {
				fmt.Println("\n Flag:[-input]/[-i] Le fichier choisi est : " + m.filepicker.Styles.Selected.Render(file) + "\n")
			}
			readFile := read(mm.selectedFile)
			strContent := string(readFile)
			fmt.Println(strContent + "\n" + "\n")
			strReplace := ask("\nPattern à remplacer:", "Choisir un pattern à remplacer est obligatoire.")
			strReplaceBy := ask("\nPattern de remplacement:", "Choisir un pattern de remplacement est obligatoire.")
			fmt.Println("\n Flag:[-replace]/[-r] Le pattern a remplacer est :" + strReplace + "\n" + "Flag:[-by]/[-b] Le pattern de remplacement est :" + strReplaceBy + "\n")
			// changeOnFile(strContent) //completer la fonction
			strContent = strings.ReplaceAll(strContent, strReplace, strReplaceBy)

			fp = filepicker.New()
			fp.DirAllowed = true
			fp.FileAllowed = false
			// pour definire le repertoire courant au repertoire user  remplacer par :
			// fp.CurrentDirectory = "./"
			fp.CurrentDirectory, _ = os.UserHomeDir()

			d := modelDirectoryPicker{
				directorypicker: fp,
			}
			td, _ := tea.NewProgram(&d).Run()
			dd := td.(modelDirectoryPicker)
			nameOfFile := ask("Quel nom de fichier voulez-vous ?", "Choisir un nom de fichier est obligatoire.")
			directory := dd.selectedDirectory + `/` + nameOfFile + ".txt"

			if file != "" {
				fmt.Println("\n Flag:[output]/[-o] Le fichier crée est : " + m.filepicker.Styles.Selected.Render(directory) + "\n")
			}

			//ecriture du nouveau fichier avec le contenu modifié
			write(directory, strContent)
			//affichage du contenu du fichier

			fmt.Print(strContent)
			fmt.Print("\n*****************************************************************************************************************************************\n")
			fmt.Print("\nPour gagner du temps à l'avenir, vous pouvez utiliser la commande ci-dessous, issue de l'action que vous venez d'effectuer, comme exemple.\n")
			fmt.Printf("martin-solving actionOnFile -i %s -o %s -r '%s' -b '%s' ", file, directory, strReplace, strReplaceBy)

		}else{
			if inputValue!="" && outputValue!="" && replaceValue!="" && replaceByValue!="" {
				readFile := read(inputValue)
				strContent := string(readFile)
				strReplace := replaceValue
				strReplaceBy := replaceByValue
				strContent = strings.ReplaceAll(strContent, strReplace, strReplaceBy)
	
				write(outputValue, strContent)
				//affichage du contenu du fichier
				fmt.Print("\n"+strContent)
				}else{
					fmt.Println("\n[Flags manquants]\nJe vous conseille de lancer la commande sans aucun flag pour lancer le mode interactif, ou simplement de compléter les flags.")
				}
			
			
			
		}
	},
}

type modelFilepicker struct {
	filepicker   filepicker.Model
	selectedFile string
	quitting     bool
	err          error
}

type modelDirectoryPicker struct {
	directorypicker   filepicker.Model
	selectedDirectory string
	quitting          bool
	err               error
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func (m modelFilepicker) Init() tea.Cmd {
	return m.filepicker.Init()
}
func (m modelDirectoryPicker) Init() tea.Cmd {
	return m.directorypicker.Init()
}

func (m modelFilepicker) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	case clearErrorMsg:
		m.err = nil
	}

	var cmd tea.Cmd
	m.filepicker, cmd = m.filepicker.Update(msg)

	// Did the user select a file?
	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
		// Get the path of the selected file.
		m.selectedFile = path
		m.quitting = true
		return m, tea.Quit
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := m.filepicker.DidSelectDisabledFile(msg); didSelect {
		// Let's clear the selectedFile and display an error.
		m.err = errors.New(path + " is not valid.")
		m.selectedFile = ""
		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	return m, cmd
}

func (m modelDirectoryPicker) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	case clearErrorMsg:
		m.err = nil
	}

	var cmd tea.Cmd
	m.directorypicker, cmd = m.directorypicker.Update(msg)

	// Did the user select a file?
	if didSelect, path := m.directorypicker.DidSelectFile(msg); didSelect {
		// Get the path of the selected file.
		m.selectedDirectory = path
		m.quitting = true
		return m, tea.Quit
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := m.directorypicker.DidSelectDisabledFile(msg); didSelect {
		// Let's clear the selectedFile and display an error.
		m.err = errors.New(path + " is not valid.")
		m.selectedDirectory = ""
		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	return m, cmd
}

func (m modelFilepicker) View() string {
	if m.quitting {
		return ""
	}
	var s strings.Builder
	s.WriteString("\n  ")
	if m.err != nil {
		s.WriteString(m.filepicker.Styles.DisabledFile.Render(m.err.Error()))
	} else if m.selectedFile == "" {
		s.WriteString("Choisissez le fichier.txt à partir duquel nous allons créer notre nouveau fichier:")
	} else {
		s.WriteString("Selected file: " + m.filepicker.Styles.Selected.Render(m.selectedFile))
	}
	s.WriteString("\n\n" + m.filepicker.View() + "\n")
	return s.String()
}

func (m modelDirectoryPicker) View() string {
	if m.quitting {
		return ""
	}
	var s strings.Builder
	s.WriteString("\n  ")
	if m.err != nil {
		s.WriteString(m.directorypicker.Styles.DisabledFile.Render(m.err.Error()))
	} else if m.selectedDirectory == "" {
		s.WriteString("Choisissez le répertoire de destination pour le nouveau fichier:")
	} else {
		s.WriteString("Selected file: " + m.directorypicker.Styles.Selected.Render(m.selectedDirectory))
	}
	s.WriteString("\n\n" + m.directorypicker.View() + "\n")
	return s.String()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func write(NewfileName string, text string) {
	if err := os.WriteFile(NewfileName, []byte(text), 0666); err != nil {
		panic(err)
	}
}

func read(filename string) string {
	data, err := os.ReadFile(filename)
	check(err)
	return string(data)
}

func ask(question string, errormsg string) string {

	var response string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(question)
	scanner.Scan()
	response = strings.Trim(scanner.Text(), " ")
	if response != "" {
		return response

	} else {
		log.Println(errormsg)
		response = ask(question, errormsg)

	}
	return response
}

// func changeOnFile(content string) string{
//     // mettre les actions de modification voulu a faire sur le fichier
// 	return content
// }

func init() {
	rootCmd.AddCommand(actionOnFileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// actionOnFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	actionOnFileCmd.Flags().StringVarP(&inputValue, "input", "i", "", "Entrez le chemin et le nom du fichier à modifier, par exemple : chemin/nom_du_fichier.txt")
	actionOnFileCmd.Flags().StringVarP(&outputValue, "output", "o", "", "Indiquez le chemin et le nom du fichier à créer en sortie, par exemple : chemin/nom_du_fichier.txt")
	actionOnFileCmd.Flags().StringVarP(&replaceValue, "replace", "r", "", "Veuillez indiquer le mot ou la série de mots à remplacer, par exemple : 'Génie' ")
	actionOnFileCmd.Flags().StringVarP(&replaceByValue, "by", "b", "", "Veuillez indiquer le mot ou la série de mots à utiliser comme remplacement, par exemple : 'Lachignol'")
}
