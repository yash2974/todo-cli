/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"TODOCLI/services"
	"github.com/spf13/cobra"
)

// fapsCmd represents the faps command
var fapsCmd = &cobra.Command{
	Use:   "faps",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		reset, _ := cmd.Flags().GetBool("reset")
		increment, _ := cmd.Flags().GetBool("inc")
		services.Faps(reset, increment)
	},
}

func init() {
	rootCmd.AddCommand(fapsCmd)
	fapsCmd.Flags().BoolP("reset", "r", false, "Reset?")
	fapsCmd.Flags().BoolP("inc", "i", false, "Increment")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fapsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fapsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
