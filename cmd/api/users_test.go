package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Intro0/goSocial/internal/store"
	"github.com/stretchr/testify/mock"
)

func TestGetUser(t *testing.T) {
	app, _ := newTestApplication(t, config{})
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
		app, mockCacheStore := newTestApplication(t, config{
			redis: redisConfig{enabled: true},
		})
		mux := app.mount()
		cachedUser := &store.User{ID: 42}

		mockCacheStore.On("Get", int64(1)).Return(nil, nil).Once()
		mockCacheStore.On("Set", mock.Anything).Return(nil).Once()
		mockCacheStore.On("Get", int64(42)).Return(cachedUser, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/v1/users/42/", nil)
		req.Header.Set("Authorization", "Bearer "+testToken)
		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)

		var response struct {
			Data store.User `json:"data"`
		}
		if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
			t.Fatal(err)
		}
		if response.Data.ID != cachedUser.ID {
			t.Errorf("response user ID = %d, want %d", response.Data.ID, cachedUser.ID)
		}

		mockCacheStore.AssertExpectations(t)
	})

	t.Run("does not use the cache when Redis is disabled", func(t *testing.T) {
		app, mockCacheStore := newTestApplication(t, config{})
		mux := app.mount()

		req := httptest.NewRequest(http.MethodGet, "/v1/users/1/", nil)
		req.Header.Set("Authorization", "Bearer "+testToken)
		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)
		mockCacheStore.AssertNotCalled(t, "Get")
		mockCacheStore.AssertNotCalled(t, "Set")
	})
}
