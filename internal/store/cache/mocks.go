package cache

import (
	"context"

	"github.com/Intro0/goSocial/internal/store"
	"github.com/stretchr/testify/mock"
)

func NewMockStorage() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

type MockUserStore struct {
	mock.Mock
}

func (m *MockUserStore) Get(_ context.Context, userID int64) (*store.User, error) {
	args := m.Called(userID)
	user, _ := args.Get(0).(*store.User)

	return user, args.Error(1)
}

func (m *MockUserStore) Set(_ context.Context, user *store.User) error {
	args := m.Called(user)

	return args.Error(0)
}
