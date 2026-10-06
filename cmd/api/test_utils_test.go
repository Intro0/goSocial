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

func newTestApplication(t *testing.T, cfg config) (*application, *cache.MockUserStore) {
	t.Helper()

	mockCacheStore := cache.NewMockUserStore()
	app := &application{
		config:        cfg,
		store:         store.NewMockStorage(),
		cacheStorage:  cache.Storage{Users: mockCacheStore},
		logger:        zap.NewNop().Sugar(),
		authenticator: &auth.TestAuthenticator{},
	}

	return app, mockCacheStore
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
