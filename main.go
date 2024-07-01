/*
Copyright © 2024 Lachignol <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Lachignol/cli-app/cmd"
	"github.com/Lachignol/cli-app/database"
)

func main() {
	database.OpenDb()
	cmd.Execute()
}
