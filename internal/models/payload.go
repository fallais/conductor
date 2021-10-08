package models

import (
	"net"

	validation "github.com/go-ozzo/ozzo-validation"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Message is a log message.
type Message struct {
	SourceIP      net.IP
	DestinationIP net.IP
	Payload       string
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewMessage returns a new Message.
func NewMessage(name, description string, value string) (*Message, error) {
	// Create the model
	model := &Message{}

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
func (model *Message) Validate() error {
	return validation.ValidateStruct(model) //validation.Field(&model.ID, validation.Required),
	//validation.Field(&model.Name, validation.Required),

}
