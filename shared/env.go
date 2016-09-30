package shared

// Step ...
type Step struct {
	Device    string   `yaml:"device" json:"device"`
	Events    int      `yaml:"events" json:"events"`
	Identifer string   `yaml:"identifier" json:"identifier"`
	Src       string   `yaml:"src" json:"src"`
	Dst       string   `yaml:"dst" json:"dst"`
	Users     []string `yaml:"users" json:"users"`
}

// Scenario ..
type Scenario struct {
	ID    int             `yaml:"id" json:"id"`
	Name  string          `yaml:"name" json:"name"`
	Steps map[string]Step `yaml:"steps" json:"steps"`
}

// Scenarii ...
var Scenarii map[string]Scenario
