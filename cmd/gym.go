/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"TODOCLI/services"
	"github.com/spf13/cobra"
)

// gymCmd represents the gym command
var gymCmd = &cobra.Command{
	Use:   "gym",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log, _ := cmd.Flags().GetBool("log")
		day, _ := cmd.Flags().GetString("day")
		services.TagGym(log, day)
	},
}

func init() {
	rootCmd.AddCommand(gymCmd)
	gymCmd.Flags().BoolP("log", "l", true, "Gym log")
	gymCmd.Flags().StringP("day", "d", "", "Tag day")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gymCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gymCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
