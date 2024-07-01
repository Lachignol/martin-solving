/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Lachignol/cli-app/database"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Ajouter une nouvelle note",
	Long:  `Ajouter une nouvelle note`,
	Run: func(cmd *cobra.Command, args []string) {
        nameSend := args[0]
		descriptionSend := args[1]
		newNote := database.Note{
			Name:        nameSend,
			Description: descriptionSend,
		}
		database.AddNote(newNote)
	},
}

func init() {
	noteCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
