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
	scenarii, err := ioutil.ReadFile("scenarii.yml")
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

	// Routes for API
	goji.Get("/api", routes.Hello)
	goji.Get("/api/v1/scenario/:id", scenarioCtrl.Get)
	goji.Get("/api/v1/scenario", scenarioCtrl.List)

	// Set the Goji server
	flag.Set("bind", ":5000")
	goji.Serve()
}
