package models

import (
	"github.com/dchest/uniuri"
	validation "github.com/go-ozzo/ozzo-validation"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Payload is a SOC payload.
type Payload struct {
	ID          string `json:"_id" bson:"_id" mapstructure:"_id"`
	Name        string `json:"name" bson:"name" mapstructure:"name"`
	Description string `json:"description" bson:"description" mapstructure:"description"`
	Value       string
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewPayload returns a new Payload.
func NewPayload(name, description string, value string) (*Payload, error) {
	// Create the model
	model := &Payload{
		ID:          uniuri.New(),
		Name:        name,
		Description: description,
		Value:       value,
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
func (model *Payload) Validate() error {
	return validation.ValidateStruct(model,
		validation.Field(&model.ID, validation.Required),
		validation.Field(&model.Name, validation.Required),
		validation.Field(&model.Value, validation.Required),
	)
}
