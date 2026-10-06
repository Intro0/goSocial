package cache

import (
	"context"

	"github.com/Intro0/goSocial/internal/store"
	"github.com/go-redis/redis/v8"
)

type Storage struct {
	Users interface {
		Get(context.Context, int64) (*store.User, error)
		Set(context.Context, *store.User) error
	}
}

func NewStorage(client *redis.Client) Storage {
	return Storage{
		Users: &UserStore{client: client},
	}
}
