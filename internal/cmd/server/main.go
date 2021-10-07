package server

import (
	"bytes"

	"github.com/fallais/conductor/internal/cmd/server/routes"
	"github.com/fallais/conductor/shared"

	"flag"
	"io/ioutil"
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

	// Read the config file
	scenarii, err := ioutil.ReadFile(*configurationFile)
	if err != nil {
		logrus.Fatalln(err)
	}

	// Unmarshal the configuration file
	err = yaml.Unmarshal(scenarii, &shared.Scenarii)
	if err != nil {
		logrus.Fatalln("Error while unmarshalling the scenarii: %v", err)
	}

	// Controllers
	scenarioCtrl := routes.NewScenarioController()

	// CORS Handler
	corsHandler := cors.New(cors.Options{AllowedOrigins: []string{"*"}, AllowCredentials: true, AllowedMethods: []string{"POST", "GET", "DELETE"}})
	goji.Use(corsHandler.Handler)

	// Routes for API
	goji.Get("/api", routes.Hello)
	goji.Get("/api/v1/scenario/:id", scenarioCtrl.Get)
	goji.Get("/api/v1/scenario", scenarioCtrl.List)

	// Set the Goji server
	flag.Set("bind", ":5001")
	goji.Serve()
}
