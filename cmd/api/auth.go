package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/Intro0/goSocial/internal/store"
	"github.com/google/uuid"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

type UserWithToken struct {
	*store.User
	Token string `json:"token"`
}

// registerUserHandler godoc
//
// @Summary Register a user
// @Description Create an inactive user and return an activation token.
// @Tags authentication
// @Accept json
// @Produce json
// @Param payload body RegisterUserPayload true "User credentials"
// @Success 201 {object} map[string]UserWithToken
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /authentication/user [post]
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
	}

	if err := user.Password.Set(payload.Password); err != nil {
		app.internalServiceError(w, r, err)
		return
	}

	plainToken := uuid.NewString()
	hash := sha256.Sum256([]byte(plainToken))
	hashedToken := hex.EncodeToString(hash[:])

	err := app.store.Users.CreateAndInvite(r.Context(), user, hashedToken, app.config.mail.exp)
	if err != nil {
		switch err {
		case store.ErrDuplicateEmail, store.ErrDuplicateUsername:
			app.badRequestResponse(w, r, err)
		default:
			app.internalServiceError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, UserWithToken{
		User:  user,
		Token: plainToken,
	}); err != nil {
		app.internalServiceError(w, r, err)
	}
}
