package routes

import (
	"net/http"

	"github.com/fallais/conductor/internal/utils"
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
func (ctrl *ScenarioController) List(w http.ResponseWriter, r *http.Request) {
	utils.JSONResponse(w, http.StatusOK, nil)
}

// Get a Message
func (ctrl *ScenarioController) Get(w http.ResponseWriter, r *http.Request) {
	// Retrieve the ID

	// Play all scenarii

	// Pubish the response
	utils.JSONResponse(w, http.StatusOK, nil)
}
