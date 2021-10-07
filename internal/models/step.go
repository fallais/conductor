package models

import (
	"github.com/dchest/uniuri"
	validation "github.com/go-ozzo/ozzo-validation"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Step is a scenario step.
type Step struct {
	ID          string   `json:"_id" bson:"_id" mapstructure:"_id"`
	Name        string   `json:"name" bson:"name" mapstructure:"name"`
	Description string   `json:"description" bson:"description" mapstructure:"description"`
	Events      []*Event `json:"events" bson:"events" mapstructure:"events"`
	IsEnabled   bool     `json:"is_enabled" bson:"is_enabled" mapstructure:"is_enabled"`
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewStep returns a new Step.
func NewStep(name, description string, events *Event, isEnabled bool) (*Step, error) {
	// Create the model
	model := &Step{
		ID:          uniuri.New(),
		Name:        name,
		Description: description,
		Events:      events,
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
func (model *Step) Validate() error {
	return validation.ValidateStruct(model,
		validation.Field(&model.ID, validation.Required),
		validation.Field(&model.Name, validation.Required),
	)
}
