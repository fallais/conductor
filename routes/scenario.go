package routes

import (
	"fmt"
	"net"
	"net/http"

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
		for key, step := range scenario {
			logrus.Infoln("Playing the step :", key)
			err := playStep(step)
			if err != nil {
				logrus.Errorln("error: %v", err)
			}
		}
	}

	// Pubish the response
	utils.JSONResponse(w, http.StatusOK, nil)
}

func playStep(step shared.Step) error {
	// Prepare addresses
	siemAddr, err := net.ResolveUDPAddr("udp", "192.168.7.10:514")
	senderAddr, err := net.ResolveUDPAddr("udp", step.Identifer)

	// Open the connection
	conn, err := net.DialUDP("udp", senderAddr, siemAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Send the logs
	for i := 0; i < step.Events; i++ {
		conn.Write([]byte(fmt.Sprintf("Sep 18 10:28:52 %s [job][20310]: Log message", step.Identifer)))
	}

	return nil
}
