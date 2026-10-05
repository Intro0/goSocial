package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Intro0/goSocial/internal/mailer"
	"github.com/Intro0/goSocial/internal/store"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

type CreateUserTokenPayload struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

type UserWithToken struct {
	*store.User
	Token string `json:"token"`
}

type UserInvitationTemplateData struct {
	Username      string
	ActivationURL string
}

func (app *application) userInvitationTemplateData(user *store.User, token string) UserInvitationTemplateData {
	return UserInvitationTemplateData{
		Username:      user.Username,
		ActivationURL: fmt.Sprintf("%s/confirm/%s", app.config.frontendURL, token),
	}
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
		Role: store.Role{
			Name: "user",
		},
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

	templateData := app.userInvitationTemplateData(user, plainToken)
	isSandbox := app.config.env != "production"
	status, err := app.mailer.Send(
		mailer.UserWelcomeTemplate,
		user.Username,
		user.Email,
		templateData,
		isSandbox,
	)
	if err != nil {
		app.logger.Errorw("send welcome email", "error", err)

		if deleteErr := app.store.Users.Delete(r.Context(), user.ID); deleteErr != nil {
			app.logger.Errorw("delete user after email failure", "user_id", user.ID, "error", deleteErr)
		}

		app.internalServiceError(w, r, err)
		return
	}

	app.logger.Infow("welcome email sent", "status", status, "user_id", user.ID)

	if err := app.jsonResponse(w, http.StatusCreated, UserWithToken{
		User:  user,
		Token: plainToken,
	}); err != nil {
		app.internalServiceError(w, r, err)
	}
}

// createTokenHandler godoc
//
// @Summary Create an access token
// @Description Authenticate an active user and return a signed JWT.
// @Tags authentication
// @Accept json
// @Produce json
// @Param payload body CreateUserTokenPayload true "User credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /authentication/token [post]
func (app *application) createTokenHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserTokenPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user, err := app.store.Users.GetByEmail(r.Context(), payload.Email)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.unauthorizedErrorResponse(w, r, err)
		default:
			app.internalServiceError(w, r, err)
		}
		return
	}

	if err := user.Password.Compare(payload.Password); err != nil {
		app.unauthorizedErrorResponse(w, r, err)
		return
	}

	now := time.Now()
	claims := jwt.MapClaims{
		"sub": strconv.FormatInt(user.ID, 10),
		"exp": now.Add(app.config.auth.token.exp).Unix(),
		"iat": now.Unix(),
		"nbf": now.Unix(),
		"iss": app.config.auth.token.issuer,
		"aud": app.config.auth.token.issuer,
	}

	token, err := app.authenticator.GenerateToken(claims)
	if err != nil {
		app.internalServiceError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, token); err != nil {
		app.internalServiceError(w, r, err)
	}
}
