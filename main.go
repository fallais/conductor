package main

import (
	"fmt"
	"io/ioutil"
	"net"

	"gopkg.in/yaml.v2"

	"github.com/Sirupsen/logrus"
)

// Config ..
type Config map[string]Scenario

// Scenario ..
type Scenario map[string]Step

// Step ...
type Step struct {
	Device    string   `yaml:"device"`
	Events    int      `yaml:"events"`
	Identifer string   `yaml:"identifier"`
	Src       string   `yaml:"src"`
	Dst       string   `yaml:"dst"`
	Users     []string `yaml:"users"`
}

func playStep(step Step) error {
	// Prepare addresses
	siemAddr, err := net.ResolveUDPAddr("udp", "192.168.7.10:514")
	senderAddr, err := net.ResolveUDPAddr("udp", step.Identifer)

	// Open the connection
	conn, err := net.DialUDP("udp", senderAddr, siemAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Send the logs
	for i := 0; i < step.Events; i++ {
		conn.Write([]byte(fmt.Sprintf("Sep 18 10:28:52 %s [job][20310]: Log message", step.Identifer)))
	}

	return nil
}

func main() {
	// Read the config file
	data, err := ioutil.ReadFile("configuration.yml")
	if err != nil {
		logrus.Fatalln(err)
	}

	// Unmarshal the configuration file
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logrus.Fatalln("error: %v", err)
	}

	// Play all scenarii
	for key, scenario := range config {
		logrus.Infoln("Playing the scenario :", key)
		for key, step := range scenario {
			logrus.Infoln("Playing the step :", key)
			err := playStep(step)
			if err != nil {
				logrus.Errorln("error: %v", err)
			}
		}
	}

}
