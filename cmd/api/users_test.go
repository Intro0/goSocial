package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Intro0/goSocial/internal/store"
	"github.com/Intro0/goSocial/internal/store/cache"
	"github.com/stretchr/testify/mock"
)

func TestGetUser(t *testing.T) {
	app := newTestApplication(t, config{})
	mux := app.mount()

	testToken, err := app.authenticator.GenerateToken(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("rejects an unauthenticated request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/users/1/", nil)
		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("allows an authenticated request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/users/1/", nil)
		req.Header.Set("Authorization", "Bearer "+testToken)
		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)
	})

	t.Run("uses the cache when Redis is enabled", func(t *testing.T) {
		app := newTestApplication(t, config{
			redis: redisConfig{enabled: true},
		})
		mux := app.mount()
		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)
		cachedUser := &store.User{ID: 1}

		mockCacheStore.On("Get", int64(1)).Return(nil, nil).Once()
		mockCacheStore.On("Set", mock.Anything).Return(nil).Once()
		mockCacheStore.On("Get", int64(1)).Return(cachedUser, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/v1/users/1/", nil)
		req.Header.Set("Authorization", "Bearer "+testToken)
		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)
		mockCacheStore.AssertExpectations(t)
	})

	t.Run("does not use the cache when Redis is disabled", func(t *testing.T) {
		app := newTestApplication(t, config{})
		mux := app.mount()
		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)

		req := httptest.NewRequest(http.MethodGet, "/v1/users/1/", nil)
		req.Header.Set("Authorization", "Bearer "+testToken)
		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)
		mockCacheStore.AssertNotCalled(t, "Get")
		mockCacheStore.AssertNotCalled(t, "Set")
	})
}
