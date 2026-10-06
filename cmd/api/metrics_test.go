package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMetricsEndpoint(t *testing.T) {
	publishMetrics(&sql.DB{})

	cfg := config{
		auth: authConfig{
			basic: basicConfig{
				user: "test-user",
				pass: "test-password",
			},
		},
	}
	app, _ := newTestApplication(t, cfg)
	mux := app.mount()

	t.Run("rejects an unauthenticated request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/debug/vars", nil)
		response := executeRequest(req, mux)

		checkResponseCode(t, http.StatusUnauthorized, response.Code)
	})

	t.Run("allows an authenticated request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/debug/vars", nil)
		req.SetBasicAuth(cfg.auth.basic.user, cfg.auth.basic.pass)
		response := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, response.Code)

		var metrics map[string]json.RawMessage
		if err := json.NewDecoder(response.Body).Decode(&metrics); err != nil {
			t.Fatal(err)
		}

		for _, metric := range []string{"version", "database", "goroutines"} {
			if _, ok := metrics[metric]; !ok {
				t.Errorf("metrics response is missing %q", metric)
			}
		}
	})
}
