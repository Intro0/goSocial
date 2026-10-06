package main

import (
	"database/sql"
	"expvar"
	"runtime"
)

func publishMetrics(db *sql.DB) {
	expvar.NewString("version").Set(version)
	expvar.Publish("database", expvar.Func(func() any {
		return db.Stats()
	}))
	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))
}
