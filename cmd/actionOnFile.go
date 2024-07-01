/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var printValue bool

// actionPageCmd represents the actionPage command
var actionOnFileCmd = &cobra.Command{
	Use:   "actionOnFile [fichier a lire] [fichier a crée]",
	Short: "Permet de lire le fichier en entrée et crée un fichier en sortie",
	Long:  `rentré en premier argument le nom du fichier a lire et appliquer la modification dans un nouveau fichier rentré en deuxieme argument`,
	Args:  cobra.MinimumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		//lecture du fichier
		readFile := read(args[0])
		strContent := string(readFile)
		//action sur le fichier
		// changeOnFile(strContent) //completer la fonction
		strContent = strings.ReplaceAll(strContent, "t", "T")
		//ecriture du nouveau fichier avec le contenu modifié
		write(args[1], strContent)
		//affichage du contenu du fichier
		if printValue {
			fmt.Print(strContent)
		}
	},
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
	actionOnFileCmd.Flags().BoolVarP(&printValue, "print", "p", false, "Affiche le fichier modifié")
}
