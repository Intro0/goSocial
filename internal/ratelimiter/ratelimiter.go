package ratelimiter

import "time"

type Limiter interface {
	Allow(clientID string) (allowed bool, retryAfter time.Duration)
}

type Config struct {
	RequestsPerTimeFrame int
	TimeFrame            time.Duration
	Enabled              bool
}
