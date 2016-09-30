package shared

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

// Scenarii ...
var Scenarii map[string]Scenario
