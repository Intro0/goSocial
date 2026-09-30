package store

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail    = errors.New("a user with that email already exists")
	ErrDuplicateUsername = errors.New("a user with that username already exists")
)

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  password `json:"-"`
	CreatedAt string   `json:"created_at"`
	IsActive  bool     `json:"is_active"`
}

type password struct {
	hash []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.hash = hash

	return nil
}

type UserStore struct {
	db *sql.DB
}

func (u *UserStore) Create(ctx context.Context, tx *sql.Tx, user *User) error {
	query := `
		INSERT INTO users (username, password, email)
		VALUES ($1, $2, $3) RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := tx.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password.hash,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			switch pqErr.Constraint {
			case "users_email_key":
				return ErrDuplicateEmail
			case "users_username_key":
				return ErrDuplicateUsername
			}
		}

		return err
	}

	return nil
}

func (u *UserStore) GetByID(ctx context.Context, userID int64) (*User, error) {
	query := `
		SELECT id, username, email, created_at
		FROM users
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &User{}
	err := u.db.QueryRowContext(
		ctx,
		query,
		userID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}

func (u *UserStore) CreateAndInvite(ctx context.Context, user *User, token string, invitationExpiry time.Duration) error {
	return withTx(u.db, ctx, func(tx *sql.Tx) error {
		if err := u.Create(ctx, tx, user); err != nil {
			return err
		}

		return u.createInvitation(ctx, tx, token, invitationExpiry, user.ID)
	})
}

func (u *UserStore) createInvitation(ctx context.Context, tx *sql.Tx, token string, expiry time.Duration, userID int64) error {
	query := `INSERT INTO user_invitations (token, user_id, expiry) VALUES ($1, $2, $3)`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, token, userID, time.Now().Add(expiry))
	return err
}

func (u *UserStore) Activate(ctx context.Context, token string) error {
	return withTx(u.db, ctx, func(tx *sql.Tx) error {
		user, err := u.getUserFromInvitation(ctx, tx, token)
		if err != nil {
			return err
		}

		if err := u.activateUser(ctx, tx, user.ID); err != nil {
			return err
		}

		return u.deleteInvitations(ctx, tx, user.ID)
	})
}

func (u *UserStore) getUserFromInvitation(ctx context.Context, tx *sql.Tx, token string) (*User, error) {
	query := `
		SELECT u.id, u.username, u.email, u.created_at, u.is_active
		FROM users u
		JOIN user_invitations ui ON u.id = ui.user_id
		WHERE ui.token = $1 AND ui.expiry > $2
	`

	hash := sha256.Sum256([]byte(token))
	hashedToken := hex.EncodeToString(hash[:])

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &User{}
	err := tx.QueryRowContext(ctx, query, hashedToken, time.Now()).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.IsActive,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return user, nil
}

func (u *UserStore) activateUser(ctx context.Context, tx *sql.Tx, userID int64) error {
	query := `UPDATE users SET is_active = TRUE WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, userID)
	return err
}

func (u *UserStore) deleteInvitations(ctx context.Context, tx *sql.Tx, userID int64) error {
	query := `DELETE FROM user_invitations WHERE user_id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, userID)
	return err
}

func (u *UserStore) Delete(ctx context.Context, userID int64) error {
	return withTx(u.db, ctx, func(tx *sql.Tx) error {
		if err := u.deleteInvitations(ctx, tx, userID); err != nil {
			return err
		}

		return u.delete(ctx, tx, userID)
	})
}

func (u *UserStore) delete(ctx context.Context, tx *sql.Tx, userID int64) error {
	query := `DELETE FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, userID)
	return err
}
