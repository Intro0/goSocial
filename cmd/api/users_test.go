package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	app := newTestApplication(t, config{})
	mux := app.mount()

	testToken, err := app.authenticator.GenerateToken(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("rejects an unauthenticated request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/users/1/", nil)
		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("allows an authenticated request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/users/1/", nil)
		req.Header.Set("Authorization", "Bearer "+testToken)
		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)
	})
}
