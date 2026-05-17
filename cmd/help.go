/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "commands",
	Long: `Commands for todocli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
		================ TODO CLI ================

		Commands:

		list
			Show all tasks

		add
			Create a new task

		detail <task_id>
			Show details of a specific task

		help
			Show available commands

		exit
			Close the application

		==========================================
		`)
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
