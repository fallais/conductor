package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "conductor",
	Short: "Send fake data to log collectors",
	Long:  "Send fake data (syslog) to log collectors such as SIEM",
}

func init() {
	rootCmd.AddCommand(playCmd)
	rootCmd.AddCommand(serverCmd)
}

// Execute the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
