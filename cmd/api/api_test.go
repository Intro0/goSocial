package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
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
				if expectedStatus == http.StatusTooManyRequests {
					retryAfter, err := strconv.Atoi(response.Header.Get("Retry-After"))
					if err != nil {
						t.Errorf("request %d Retry-After = %q, want an integer", request+1, response.Header.Get("Retry-After"))
					}
					if retryAfter < 1 || retryAfter > int(cfg.rateLimiter.TimeFrame.Seconds()) {
						t.Errorf("request %d Retry-After = %d, want 1 through %d", request+1, retryAfter, int(cfg.rateLimiter.TimeFrame.Seconds()))
					}
				}
			}
		})
	}
}
