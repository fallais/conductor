package models

import (
	"github.com/dchest/uniuri"
	validation "github.com/go-ozzo/ozzo-validation"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Scenario is a SOC scenario.
type Scenario struct {
	ID          string  `json:"_id" bson:"_id" mapstructure:"_id"`
	Name        string  `json:"name" bson:"name" mapstructure:"name"`
	Description string  `json:"description" bson:"description" mapstructure:"description"`
	Steps       []*Step `json:"steps" bson:"steps" mapstructure:"steps"`
	IsEnabled   bool    `json:"is_enabled" bson:"is_enabled" mapstructure:"is_enabled"`
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewScenario returns a new Scenario.
func NewScenario(name, description string, steps []*Step, isEnabled bool) (*Scenario, error) {
	// Create the model
	model := &Scenario{
		ID:          uniuri.New(),
		Name:        name,
		Description: description,
		Steps:       steps,
		IsEnabled:   isEnabled,
	}

	// Validate the model
	err := model.Validate()
	if err != nil {
		return nil, err
	}

	return model, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Validate the model.
func (model *Scenario) Validate() error {
	return validation.ValidateStruct(model,
		validation.Field(&model.ID, validation.Required),
		validation.Field(&model.Name, validation.Required),
	)
}
