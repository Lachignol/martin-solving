/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/Lachignol/martin-solving/note"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Changer le status de la tache",
	Long:  `Ajouter une nouvelle note`,
	Run: func(cmd *cobra.Command, args []string) {
		number, err := strconv.Atoi(args[0])
		if err != nil {
			log.Println(err)
		}
		note.ToggleTodo(number)
	},
	}


func init() {
	noteCmd.AddCommand(toggleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}