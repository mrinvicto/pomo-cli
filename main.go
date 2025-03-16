/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/mrinvicto/pomo-cli/cmd"
	"github.com/mrinvicto/pomo-cli/src/config"
	"github.com/mrinvicto/pomo-cli/src/database"
)

func main() {
	fmt.Println("Main started")
	config.InitConfig()
	config := config.GetConfig()
	db := database.InitDatabase(config)
	db.InitDB()
	cmd.Execute()
}
