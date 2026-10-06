package main

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/Intro0/goSocial/internal/store"
	"github.com/Intro0/goSocial/internal/store/cache"
	"go.uber.org/zap"
)

type fakeUserStore struct {
	getByID func(context.Context, int64) (*store.User, error)
}

func (s *fakeUserStore) Create(context.Context, *sql.Tx, *store.User) error {
	return errors.New("not implemented")
}

func (s *fakeUserStore) CreateAndInvite(context.Context, *store.User, string, time.Duration) error {
	return errors.New("not implemented")
}

func (s *fakeUserStore) Activate(context.Context, string) error {
	return errors.New("not implemented")
}

func (s *fakeUserStore) Delete(context.Context, int64) error {
	return errors.New("not implemented")
}

func (s *fakeUserStore) GetByID(ctx context.Context, userID int64) (*store.User, error) {
	return s.getByID(ctx, userID)
}

func (s *fakeUserStore) GetByEmail(context.Context, string) (*store.User, error) {
	return nil, errors.New("not implemented")
}

type fakeCachedUserStore struct {
	get func(context.Context, int64) (*store.User, error)
	set func(context.Context, *store.User) error
}

func (s *fakeCachedUserStore) Get(ctx context.Context, userID int64) (*store.User, error) {
	return s.get(ctx, userID)
}

func (s *fakeCachedUserStore) Set(ctx context.Context, user *store.User) error {
	return s.set(ctx, user)
}

func newCacheTestApplication(enabled bool, users *fakeUserStore, cachedUsers *fakeCachedUserStore) *application {
	return &application{
		config: config{redis: redisConfig{enabled: enabled}},
		store: store.Storage{
			Users: users,
		},
		cacheStorage: cache.Storage{
			Users: cachedUsers,
		},
		logger: zap.NewNop().Sugar(),
	}
}

func TestGetUserUsesCacheHit(t *testing.T) {
	cachedUser := &store.User{ID: 1, Username: "cached"}
	app := newCacheTestApplication(
		true,
		&fakeUserStore{getByID: func(context.Context, int64) (*store.User, error) {
			t.Fatal("database should not be called on a cache hit")
			return nil, nil
		}},
		&fakeCachedUserStore{
			get: func(context.Context, int64) (*store.User, error) { return cachedUser, nil },
			set: func(context.Context, *store.User) error {
				t.Fatal("cache should not be written on a cache hit")
				return nil
			},
		},
	)

	user, err := app.getUser(t.Context(), cachedUser.ID)
	if err != nil {
		t.Fatalf("getUser returned an error: %v", err)
	}
	if user != cachedUser {
		t.Fatalf("getUser returned %+v, want %+v", user, cachedUser)
	}
}

func TestGetUserCachesDatabaseResultOnMiss(t *testing.T) {
	databaseUser := &store.User{ID: 1, Username: "database"}
	wasCached := false
	app := newCacheTestApplication(
		true,
		&fakeUserStore{getByID: func(context.Context, int64) (*store.User, error) { return databaseUser, nil }},
		&fakeCachedUserStore{
			get: func(context.Context, int64) (*store.User, error) { return nil, nil },
			set: func(_ context.Context, user *store.User) error {
				if user != databaseUser {
					t.Fatalf("cached user = %+v, want %+v", user, databaseUser)
				}
				wasCached = true
				return nil
			},
		},
	)

	user, err := app.getUser(t.Context(), databaseUser.ID)
	if err != nil {
		t.Fatalf("getUser returned an error: %v", err)
	}
	if user != databaseUser {
		t.Fatalf("getUser returned %+v, want %+v", user, databaseUser)
	}
	if !wasCached {
		t.Fatal("database user was not cached")
	}
}

func TestGetUserFallsBackWhenCacheFails(t *testing.T) {
	databaseUser := &store.User{ID: 1, Username: "database"}
	app := newCacheTestApplication(
		true,
		&fakeUserStore{getByID: func(context.Context, int64) (*store.User, error) { return databaseUser, nil }},
		&fakeCachedUserStore{
			get: func(context.Context, int64) (*store.User, error) {
				return nil, errors.New("cache unavailable")
			},
			set: func(context.Context, *store.User) error {
				return errors.New("cache unavailable")
			},
		},
	)

	user, err := app.getUser(t.Context(), databaseUser.ID)
	if err != nil {
		t.Fatalf("getUser returned an error: %v", err)
	}
	if user != databaseUser {
		t.Fatalf("getUser returned %+v, want %+v", user, databaseUser)
	}
}

func TestGetUserSkipsCacheWhenDisabled(t *testing.T) {
	databaseUser := &store.User{ID: 1, Username: "database"}
	app := newCacheTestApplication(
		false,
		&fakeUserStore{getByID: func(context.Context, int64) (*store.User, error) { return databaseUser, nil }},
		&fakeCachedUserStore{
			get: func(context.Context, int64) (*store.User, error) {
				t.Fatal("cache should not be read when disabled")
				return nil, nil
			},
			set: func(context.Context, *store.User) error {
				t.Fatal("cache should not be written when disabled")
				return nil
			},
		},
	)

	user, err := app.getUser(t.Context(), databaseUser.ID)
	if err != nil {
		t.Fatalf("getUser returned an error: %v", err)
	}
	if user != databaseUser {
		t.Fatalf("getUser returned %+v, want %+v", user, databaseUser)
	}
}
