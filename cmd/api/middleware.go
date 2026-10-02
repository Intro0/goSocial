package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func parseBasicAuth(header string) (username string, password string, err error) {
	if header == "" {
		return "", "", fmt.Errorf("authorization header is missing")
	}

	scheme, encodedCredentials, found := strings.Cut(header, " ")
	if !found || scheme != "Basic" || encodedCredentials == "" {
		return "", "", fmt.Errorf("authorization header is malformed")
	}

	decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		return "", "", err
	}

	username, password, found = strings.Cut(string(decodedCredentials), ":")
	if !found {
		return "", "", fmt.Errorf("basic credentials are malformed")
	}

	return username, password, nil
}

func (app *application) BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, err := parseBasicAuth(r.Header.Get("Authorization"))
		if err != nil {
			app.unauthorizedBasicErrorResponse(w, r, err)
			return
		}

		if username != app.config.auth.basic.user || password != app.config.auth.basic.pass {
			app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("invalid basic auth credentials"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
