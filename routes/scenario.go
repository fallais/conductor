package routes

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/zenazn/goji/web"

	"orchestrator/shared"
	"orchestrator/utils"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// ScenarioController ...
type ScenarioController struct {
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewScenarioController ...
func NewScenarioController() *ScenarioController {
	return &ScenarioController{}
}

//------------------------------------------------------------------------------
// Protocol
//------------------------------------------------------------------------------

// ControllerError contains ...
type ControllerError struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

//------------------------------------------------------------------------------
// Routes
//------------------------------------------------------------------------------

// List all Scenario
func (ctrl *ScenarioController) List(c web.C, w http.ResponseWriter, r *http.Request) {
	utils.JSONResponse(w, http.StatusOK, shared.Scenarii)
}

// Get a Message
func (ctrl *ScenarioController) Get(c web.C, w http.ResponseWriter, r *http.Request) {
	// Play all scenarii
	for key, scenario := range shared.Scenarii {
		logrus.Infoln("Playing the scenario :", key)
		for key, step := range scenario.Steps {
			logrus.Infoln("Playing the step :", key)
			err := playStep(step)
			if err != nil {
				logrus.Errorln("Error:", err)
			}
		}
	}

	// Pubish the response
	utils.JSONResponse(w, http.StatusOK, nil)
}

func playStep(step shared.Step) error {
	// Prepare addresses
	siemAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:514")
	 if err != nil {
                return err
        }
	senderAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:50000", step.Events.LogSourceIP))
	if err != nil {
                return err
       }

	// Open the connection
	conn, err := net.DialUDP("udp", senderAddr, siemAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Prepare the Payload
	payload := step.Events.Payload
	for key, value := range step.Events.Values {
		payload = strings.Replace(payload, fmt.Sprintf("{{%s}}", key), value[rand.Intn(len(value))], -1)
	}

	// Send the logs
	for i := 0; i < step.Events.Nb; i++ {
		conn.Write([]byte(step.Events.Payload))
	}

	return nil
}
