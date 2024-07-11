/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Lachignol/martin-solving/database"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Permet de supprimé proprement l'application.",
	Long:  "Supprime la base de données ainsi que son répertoire. Ensuite, il ne reste plus qu'à supprimer le binaire (instructions détaillées dans la commande).",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Voulez-vous vraiment supprimé l'application [y]/[n]? \n")
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		if response == "yes" || response == "y" {
			database.Uninstall()
		} else {
			fmt.Println("Supression de l'application annulé")
		}

	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
