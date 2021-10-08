package cmd

import (
	"github.com/fallais/conductor/internal/cmd/play"

	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play the scenario",
	Run:   play.Run,
}

func init() {
	playCmd.Flags().StringP("scenario", "s", "scenario.yaml", "Scenario YAML file")
}
