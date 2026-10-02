package main

import (
	"net/http"
)

func (app *application) internalServiceError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("internal server error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err,
	)

	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("bad request",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err,
	)

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("resource not found",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err,
	)

	writeJSONError(w, http.StatusNotFound, "not found")
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("resource conflict",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err,
	)

	writeJSONError(w, http.StatusConflict, "resource already exists")
}

func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("unauthorized request",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err,
	)

	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *application) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("unauthorized basic auth request",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err,
	)

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}
