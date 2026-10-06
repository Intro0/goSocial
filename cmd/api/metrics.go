package main

import (
	"database/sql"
	"expvar"
	"runtime"
	"sync"
)

var publishMetricsOnce sync.Once

func publishMetrics(db *sql.DB) {
	publishMetricsOnce.Do(func() {
		expvar.NewString("version").Set(version)
		expvar.Publish("database", expvar.Func(func() any {
			return db.Stats()
		}))
		expvar.Publish("goroutines", expvar.Func(func() any {
			return runtime.NumGoroutine()
		}))
	})
}
