package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "simple-reader",
	Short: "File Reader (Support JSON)",
	Long:  `A tool help you read file data with key name and return data which is friendly with command line`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
