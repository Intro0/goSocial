package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

func NewMockStorage() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

type MockUserStore struct{}

func (m *MockUserStore) Create(context.Context, *sql.Tx, *User) error {
	return errors.New("not implemented")
}

func (m *MockUserStore) CreateAndInvite(context.Context, *User, string, time.Duration) error {
	return errors.New("not implemented")
}

func (m *MockUserStore) Activate(context.Context, string) error {
	return errors.New("not implemented")
}

func (m *MockUserStore) Delete(context.Context, int64) error {
	return errors.New("not implemented")
}

func (m *MockUserStore) GetByID(_ context.Context, userID int64) (*User, error) {
	return &User{ID: userID}, nil
}

func (m *MockUserStore) GetByEmail(context.Context, string) (*User, error) {
	return nil, errors.New("not implemented")
}
