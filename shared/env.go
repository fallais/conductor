package shared

// Events ...
type Events struct {
	Nb          int                 `yaml:"nb" json:"nb"`
	LogSourceIP string              `yaml:"log_source_ip" json:"log_source_ip"`
	Payload     string              `yaml:"payload" json:"payload"`
	Behavior    string              `yaml:"behavior" json:"behavior"`
	Values      map[string][]string `yaml:"values" json:"values"`
}

// Step ...
type Step struct {
	Name   string `yaml:"name" json:"name"`
	Events Events `yaml:"events" json:"events"`
}

// Scenario ..
type Scenario struct {
	ID    int             `yaml:"id" json:"id"`
	Name  string          `yaml:"name" json:"name"`
	Steps map[string]Step `yaml:"steps" json:"steps"`
}

// Scenarii ...
var Scenarii map[string]Scenario
