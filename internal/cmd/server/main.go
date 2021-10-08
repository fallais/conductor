package server

import (
	"bytes"
	"flag"
	"io/ioutil"

	"github.com/fallais/conductor/internal/cmd/server/routes"

	"github.com/goji/goji"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Run is a convenient function for Cobra.
func Run(cmd *cobra.Command, args []string) {
	// Flag
	configFile, err := cmd.Flags().GetString("config")
	if err != nil {
		logrus.WithError(err).Fatalln("Error with the configuration file flag")
	}

	// Read configuration file
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		logrus.WithError(err).Fatalln("Error when reading configuration file")
	}

	// Initialize values with Viper
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		logrus.WithError(err).Fatalln("Error when reading configuration data")
	}

	// Controllers
	scenarioCtrl := routes.NewScenarioController()

	// Routes for API
	goji.Get("/api", routes.Hello)
	goji.Get("/api/v1/scenario/:id", scenarioCtrl.Get)
	goji.Get("/api/v1/scenario", scenarioCtrl.List)

	// Set the Goji server
	flag.Set("bind", ":5001")
	goji.Serve()
}
