package simple

import (
	"context"
	"time"

	"conductor/internal/generators"
	"conductor/internal/models"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// simpleGenerator is the simplest event generator.
// It takes a template a repeat it as many times as the count specified.
type simpleGenerator struct {
	template string
	count    int
	latency  time.Duration
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// New return an Event.
func New(template string, count int) (generators.Generator, error) {
	return &simpleGenerator{
		template: template,
		count:    count,
		latency:  1 * time.Second,
	}, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Generate the events into a channel.
func (g *simpleGenerator) Generate(ctx context.Context) <-chan *models.Payload {
	// Create the channel
	payloadCh := make(chan *models.Payload)

	go func() {
	loop:
		for i := 0; i < g.count; i++ {
			// Stop if the context is done
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			// Create the payload
			payload := &models.Payload{
				Value: g.template,
			}

			// Append
			payloadCh <- payload

			// Apply the latency
			time.Sleep(g.latency)
		}

		// Close the channel
		close(payloadCh)
	}()

	return payloadCh
}
