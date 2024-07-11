/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var todoCmd = &cobra.Command{
	Use:   "todo [command]",
	Short: "Todo-list",
	Long:  `N'importe quoi que tu veu ajouter a la liste afin de ne pas l'oublier`,
}

func init() {
	rootCmd.AddCommand(todoCmd)

}
