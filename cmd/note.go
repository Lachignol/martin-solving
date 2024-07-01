/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	

	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "N'importe quoi que tu veu ajouter a la liste afin de ne pas l'oublier",
	Long: `N'importe quoi que tu veu ajouter a la liste afin de ne pas l'oublier`,

}

func init() {
	rootCmd.AddCommand(noteCmd)


}
