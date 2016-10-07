package main

import (
	"flag"
	"io/ioutil"
	"time"

	"orchestrator/routes"
	"orchestrator/shared"

	"gopkg.in/yaml.v2"

	"github.com/Sirupsen/logrus"
	"github.com/rs/cors"
	"github.com/zenazn/goji"
)

var configurationFile = flag.String("c", "scenarii.yml", "Specify the location of the configuration file")

func main() {
	// Set localtime to UTC
	time.Local = time.UTC

	// Parse the flags
	flag.Parse()

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
