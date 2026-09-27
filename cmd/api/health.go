package main

import (
	"net/http"
)

// healthCheckHandler godoc
//
// @Summary Check API health
// @Tags ops
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}
	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServiceError(w, r, err)
	}
}
