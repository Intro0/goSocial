package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
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

func parseBearerToken(header string) (string, error) {
	if header == "" {
		return "", fmt.Errorf("authorization header is missing")
	}

	scheme, token, found := strings.Cut(header, " ")
	if !found || scheme != "Bearer" || token == "" {
		return "", fmt.Errorf("authorization header is malformed")
	}

	return token, nil
}

func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := parseBearerToken(r.Header.Get("Authorization"))
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		jwtToken, err := app.authenticator.ValidateToken(tokenString)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		claims, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("token claims have an unexpected type"))
			return
		}

		subject, err := claims.GetSubject()
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		userID, err := strconv.ParseInt(subject, 10, 64)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		user, err := app.store.Users.GetByID(r.Context(), userID)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
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
