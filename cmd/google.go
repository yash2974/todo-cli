/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"TODOCLI/google"
	"github.com/spf13/cobra"
)

// googleCmd represents the google command
var googleCmd = &cobra.Command{
	Use:   "google",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskType, _ := cmd.Flags().GetString("type")
		startTime, _ := cmd.Flags().GetString("start_time")
		endTime, _ := cmd.Flags().GetString("end_time")
		title, _ := cmd.Flags().GetString("title")
		location, _ := cmd.Flags().GetString("location")
		description, _ := cmd.Flags().GetString("description")
		timeZone, _ := cmd.Flags().GetString("timeZone")
		google.Calender(taskType, startTime, endTime, title, location, description, timeZone)
	},
}

func init() {
	rootCmd.AddCommand(googleCmd)
	googleCmd.Flags().StringP("type", "t", "", "Enter Task Type")
	googleCmd.Flags().StringP("start_time", "s", "", "Enter Start Time")
	googleCmd.Flags().StringP("end_time", "e", "", "Enter End Time")
	googleCmd.Flags().StringP("title", "T", "", "Enter Title/Summary")
	googleCmd.Flags().StringP("location", "l", "", "Enter Location")
	googleCmd.Flags().StringP("description", "d", "", "Enter Description")
	googleCmd.Flags().StringP("timeZone", "z", "Asia/Kolkata", "Enter Timezone")
	googleCmd.MarkFlagRequired("type")
	googleCmd.MarkFlagRequired("start_time")
	googleCmd.MarkFlagRequired("end_time")
	googleCmd.MarkFlagRequired("title")
	googleCmd.MarkFlagRequired("location")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// googleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// googleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
