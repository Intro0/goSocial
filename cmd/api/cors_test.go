package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORSAllowsConfiguredOrigin(t *testing.T) {
	cfg := config{
		corsAllowedOrigin: "http://localhost:5174",
	}
	app, _ := newTestApplication(t, cfg)

	req := httptest.NewRequest(http.MethodGet, "/v1/health", nil)
	req.Header.Set("Origin", cfg.corsAllowedOrigin)

	response := executeRequest(req, app.mount())
	allowedOrigin := response.Header().Get("Access-Control-Allow-Origin")
	if allowedOrigin != cfg.corsAllowedOrigin {
		t.Errorf("Access-Control-Allow-Origin = %q, want %q", allowedOrigin, cfg.corsAllowedOrigin)
	}
}

func TestCORSRejectsUnconfiguredOrigin(t *testing.T) {
	cfg := config{
		corsAllowedOrigin: "http://localhost:5174",
	}
	app, _ := newTestApplication(t, cfg)

	req := httptest.NewRequest(http.MethodGet, "/v1/health", nil)
	req.Header.Set("Origin", "https://untrusted.example.com")

	response := executeRequest(req, app.mount())
	if allowedOrigin := response.Header().Get("Access-Control-Allow-Origin"); allowedOrigin != "" {
		t.Errorf("Access-Control-Allow-Origin = %q, want empty", allowedOrigin)
	}
}
