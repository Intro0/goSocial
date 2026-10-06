package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Intro0/goSocial/internal/auth"
	"github.com/Intro0/goSocial/internal/store"
	"github.com/Intro0/goSocial/internal/store/cache"
	"go.uber.org/zap"
)

func newTestApplication(t *testing.T, cfg config) *application {
	t.Helper()

	return &application{
		config:        cfg,
		store:         store.NewMockStorage(),
		cacheStorage:  cache.NewMockStorage(),
		logger:        zap.NewNop().Sugar(),
		authenticator: &auth.TestAuthenticator{},
	}
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	t.Helper()

	if expected != actual {
		t.Errorf("response status = %d, want %d", actual, expected)
	}
}
