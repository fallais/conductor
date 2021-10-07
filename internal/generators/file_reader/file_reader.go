package file_reader

import (
	"bufio"
	"context"
	"io"
	"time"

	"conductor/internal/generators"
	"conductor/internal/models"

	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// fileReaderGenerator just reads a file and send the data within it.
type fileReaderGenerator struct {
	file    io.Reader
	latency time.Duration
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// New return a new file reader generator.
func New(file io.Reader) (generators.Generator, error) {
	return &fileReaderGenerator{
		file:    file,
		latency: 1 * time.Second,
	}, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Generate the events into a channel.
func (g *fileReaderGenerator) Generate(ctx context.Context) <-chan *models.Payload {
	// Create the channel
	payloadCh := make(chan *models.Payload)

	go func() {
		scanner := bufio.NewScanner(g.file)
		for scanner.Scan() {
			// Create the payload
			payload := &models.Payload{
				Value: scanner.Text(),
			}

			// Send to channel
			payloadCh <- payload

			// Apply the latency
			time.Sleep(g.latency)
		}

		err := scanner.Err()
		if err != nil {
			logrus.WithError(err).Errorln("error while scanning the file")
		}

		// Close the channel
		close(payloadCh)
	}()

	return payloadCh
}
