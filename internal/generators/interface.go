package generators

import (
	"context"

	"conductor/internal/models"
)

// Generator is an payload generator.
type Generator interface {
	Generate(context.Context) <-chan *models.Payload
}
