/*
Copyright © 2024 Lachignol <scordilisalexandre7@gmail.com>
*/
package main

import (
	_ "embed"

	"github.com/Lachignol/martin-solving/cmd"
	"github.com/Lachignol/martin-solving/database"
)

func main() {
	database.OpenDb()
	database.CreateTable()
	cmd.Execute()
}
