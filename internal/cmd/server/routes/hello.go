package routes

import (
	"net/http"

	"github.com/fallais/conductor/internal/utils"
)

//------------------------------------------------------------------------------
// Resources
//------------------------------------------------------------------------------

// HelloResponse contains the result of the Hello request.
type HelloResponse struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

//------------------------------------------------------------------------------
// Routes
//------------------------------------------------------------------------------

// Hello shows basic information about the API on its frontpage.
func Hello(w http.ResponseWriter, r *http.Request) {
	utils.JSONResponse(w, http.StatusOK, &HelloResponse{
		Message: "Dashboard API",
		Version: "1",
	})
}
