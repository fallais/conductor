package cmd

import (
	"github.com/fallais/conductor/internal/play"

	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play the scenario",
	Run:   play.Select,
}

func init() {
	playCmd.Flags().StringP("scenario", "s", "scenario.yaml", "Scenario YAML file")
}
