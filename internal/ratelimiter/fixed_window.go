package ratelimiter

import (
	"sync"
	"time"
)

type FixedWindowLimiter struct {
	mu          sync.Mutex
	clients     map[string]clientWindow
	limit       int
	window      time.Duration
	nextCleanup time.Time
}

type clientWindow struct {
	count   int
	resetAt time.Time
}

func NewFixedWindowLimiter(limit int, window time.Duration) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		clients: make(map[string]clientWindow),
		limit:   limit,
		window:  window,
	}
}

func (l *FixedWindowLimiter) Allow(clientID string) (bool, time.Duration) {
	now := time.Now()

	l.mu.Lock()
	defer l.mu.Unlock()

	if !now.Before(l.nextCleanup) {
		l.cleanupExpiredClients(now)
		l.nextCleanup = now.Add(l.window)
	}

	client := l.clients[clientID]
	if client.resetAt.IsZero() || !now.Before(client.resetAt) {
		client = clientWindow{resetAt: now.Add(l.window)}
	}

	if client.count >= l.limit {
		return false, client.resetAt.Sub(now)
	}

	client.count++
	l.clients[clientID] = client

	return true, 0
}

func (l *FixedWindowLimiter) cleanupExpiredClients(now time.Time) {
	for clientID, client := range l.clients {
		if !now.Before(client.resetAt) {
			delete(l.clients, clientID)
		}
	}
}
