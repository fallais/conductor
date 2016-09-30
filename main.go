package main

import (
	"flag"
	"io/ioutil"

	"orchestrator/routes"
	"orchestrator/shared"

	"gopkg.in/yaml.v2"

	"github.com/Sirupsen/logrus"
	"github.com/zenazn/goji"
)

func main() {
	// Read the config file
	data, err := ioutil.ReadFile("configuration.yml")
	if err != nil {
		logrus.Fatalln(err)
	}

	// Unmarshal the configuration file
	err = yaml.Unmarshal(data, &shared.Scenarii)
	if err != nil {
		logrus.Fatalln("error: %v", err)
	}

	scenarioCtrl := routes.NewScenarioController()

	// Routes for API
	goji.Get("/api", routes.Hello)
	goji.Get("/api/v1/scenario/:id", scenarioCtrl.Get)
	goji.Get("/api/v1/scenario", scenarioCtrl.List)

	flag.Set("bind", ":5000")
	goji.Serve()

}
