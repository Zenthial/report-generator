/*
Copyright © 2022 Thomas Schollenberger <tom@schollenbergers.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Creates the default report",
	Long: "This command creates the chapter wide report. It will generate every single position listed in positions.go.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called")
	},
}

func init() {
	rootCmd.AddCommand(eboardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rootCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Execute() error {
	return rootCmd.Execute()
}