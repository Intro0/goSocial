package ratelimiter

import (
	"testing"
	"time"
)

func newTestFixedWindowLimiter(
	limit int,
	window time.Duration,
	now func() time.Time,
) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		clients: make(map[string]clientWindow),
		limit:   limit,
		window:  window,
		now:     now,
	}
}

func TestFixedWindowLimiterAllowsRequestsUpToLimit(t *testing.T) {
	now := time.Date(2026, time.October, 6, 12, 0, 0, 0, time.UTC)
	limiter := newTestFixedWindowLimiter(2, time.Minute, func() time.Time {
		return now
	})

	for request := 1; request <= 2; request++ {
		allowed, retryAfter := limiter.Allow("client-1")
		if !allowed {
			t.Errorf("request %d was rejected, want allowed", request)
		}
		if retryAfter != 0 {
			t.Errorf("request %d retry after = %v, want 0", request, retryAfter)
		}
	}

	allowed, retryAfter := limiter.Allow("client-1")
	if allowed {
		t.Error("request above the limit was allowed")
	}
	if retryAfter <= 0 {
		t.Errorf("retry after = %v, want a positive duration", retryAfter)
	}
}

func TestFixedWindowLimiterTracksClientsIndependently(t *testing.T) {
	now := time.Date(2026, time.October, 6, 12, 0, 0, 0, time.UTC)
	limiter := newTestFixedWindowLimiter(1, time.Minute, func() time.Time {
		return now
	})

	if allowed, _ := limiter.Allow("client-1"); !allowed {
		t.Fatal("first request from client-1 was rejected")
	}
	if allowed, _ := limiter.Allow("client-1"); allowed {
		t.Fatal("second request from client-1 was allowed")
	}
	if allowed, _ := limiter.Allow("client-2"); !allowed {
		t.Fatal("first request from client-2 was rejected")
	}
}

func TestFixedWindowLimiterResetsExpiredWindows(t *testing.T) {
	now := time.Date(2026, time.October, 6, 12, 0, 0, 0, time.UTC)
	limiter := newTestFixedWindowLimiter(1, time.Minute, func() time.Time {
		return now
	})

	if allowed, _ := limiter.Allow("client-1"); !allowed {
		t.Fatal("first request was rejected")
	}
	if allowed, _ := limiter.Allow("client-1"); allowed {
		t.Fatal("request above the limit was allowed")
	}

	now = now.Add(time.Minute)

	if allowed, _ := limiter.Allow("client-1"); !allowed {
		t.Fatal("request after the window expired was rejected")
	}
}

func TestFixedWindowLimiterRemovesExpiredClients(t *testing.T) {
	now := time.Date(2026, time.October, 6, 12, 0, 0, 0, time.UTC)
	limiter := newTestFixedWindowLimiter(1, time.Minute, func() time.Time {
		return now
	})

	if allowed, _ := limiter.Allow("client-1"); !allowed {
		t.Fatal("first request from client-1 was rejected")
	}

	now = now.Add(time.Minute)

	if allowed, _ := limiter.Allow("client-2"); !allowed {
		t.Fatal("first request from client-2 was rejected")
	}

	if _, exists := limiter.clients["client-1"]; exists {
		t.Error("expired client-1 was not removed")
	}
}
