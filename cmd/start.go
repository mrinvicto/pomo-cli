/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/mrinvicto/pomo-cli/src/database"
	"github.com/mrinvicto/pomo-cli/src/helpers"
	"github.com/mrinvicto/pomo-cli/src/models"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := database.GetDatabase()
		session := models.Session{
			Title:    "Deep Work Session",
			Tags:     []string{"focus", "study"},
			Duration: 25,
			Status:   1,
		}
		db.CreateSession(&session)
		fmt.Println("Session Id", session.ID)
		helpers.Timer(20)
		exec.Command("afplay", "/System/Library/Sounds/Ping.aiff").Run()
		exec.Command("afplay", "/System/Library/Sounds/Ping.aiff").Run()
		exec.Command("afplay", "/System/Library/Sounds/Ping.aiff").Run()

	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
