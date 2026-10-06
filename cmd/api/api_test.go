package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Intro0/goSocial/internal/ratelimiter"
)

func TestRateLimiterMiddleware(t *testing.T) {
	tests := []struct {
		name     string
		enabled  bool
		limit    int
		statuses []int
	}{
		{
			name:     "rejects requests above the limit",
			enabled:  true,
			limit:    2,
			statuses: []int{http.StatusOK, http.StatusOK, http.StatusTooManyRequests},
		},
		{
			name:     "skips limits when disabled",
			enabled:  false,
			limit:    1,
			statuses: []int{http.StatusOK, http.StatusOK, http.StatusOK},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg := config{
				auth: authConfig{
					basic: basicConfig{
						user: "test-user",
						pass: "test-password",
					},
				},
				rateLimiter: ratelimiter.Config{
					RequestsPerTimeFrame: test.limit,
					TimeFrame:            time.Minute,
					Enabled:              test.enabled,
				},
			}

			app, _ := newTestApplication(t, cfg)
			server := httptest.NewServer(app.mount())
			defer server.Close()

			for request, expectedStatus := range test.statuses {
				req, err := http.NewRequest(http.MethodGet, server.URL+"/v1/health", nil)
				if err != nil {
					t.Fatal(err)
				}
				req.SetBasicAuth(cfg.auth.basic.user, cfg.auth.basic.pass)

				response, err := server.Client().Do(req)
				if err != nil {
					t.Fatal(err)
				}
				response.Body.Close()

				checkResponseCode(t, expectedStatus, response.StatusCode)
				if expectedStatus == http.StatusTooManyRequests && response.Header.Get("Retry-After") == "" {
					t.Errorf("request %d did not include Retry-After", request+1)
				}
			}
		})
	}
}
