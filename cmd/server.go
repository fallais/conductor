package cmd

import (
	"github.com/fallais/conductor/internal/server"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a server",
	Run:   server.Run,
}

func init() {
	playCmd.Flags().StringP("config", "c", "config.yml", "Config file")
}
