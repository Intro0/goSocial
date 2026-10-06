package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Intro0/goSocial/internal/store"
	"github.com/golang-jwt/jwt/v5"
)

func parseBasicAuth(header string) (username string, password string, err error) {
	if header == "" {
		return "", "", fmt.Errorf("authorization header is missing")
	}

	scheme, encodedCredentials, found := strings.Cut(header, " ")
	if !found || scheme != "Basic" || encodedCredentials == "" {
		return "", "", fmt.Errorf("authorization header is malformed")
	}

	decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		return "", "", err
	}

	username, password, found = strings.Cut(string(decodedCredentials), ":")
	if !found {
		return "", "", fmt.Errorf("basic credentials are malformed")
	}

	return username, password, nil
}

func parseBearerToken(header string) (string, error) {
	if header == "" {
		return "", fmt.Errorf("authorization header is missing")
	}

	scheme, token, found := strings.Cut(header, " ")
	if !found || scheme != "Bearer" || token == "" {
		return "", fmt.Errorf("authorization header is malformed")
	}

	return token, nil
}

func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := parseBearerToken(r.Header.Get("Authorization"))
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		jwtToken, err := app.authenticator.ValidateToken(tokenString)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		claims, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("token claims have an unexpected type"))
			return
		}

		subject, err := claims.GetSubject()
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		userID, err := strconv.ParseInt(subject, 10, 64)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		user, err := app.getUser(r.Context(), userID)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) getUser(ctx context.Context, userID int64) (*store.User, error) {
	if !app.config.redis.enabled {
		return app.store.Users.GetByID(ctx, userID)
	}

	user, err := app.cacheStorage.Users.Get(ctx, userID)
	if err != nil {
		app.logger.Warnw("get user from cache", "user_id", userID, "error", err)
		user = nil
	}

	if user == nil {
		user, err = app.store.Users.GetByID(ctx, userID)
		if err != nil {
			return nil, err
		}

		if err := app.cacheStorage.Users.Set(ctx, user); err != nil {
			app.logger.Warnw("cache user", "user_id", userID, "error", err)
		}
	}

	return user, nil
}

func (app *application) BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, err := parseBasicAuth(r.Header.Get("Authorization"))
		if err != nil {
			app.unauthorizedBasicErrorResponse(w, r, err)
			return
		}

		if username != app.config.auth.basic.user || password != app.config.auth.basic.pass {
			app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("invalid basic auth credentials"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) checkRolePrecedence(ctx context.Context, user *store.User, requiredRole string) (bool, error) {
	role, err := app.store.Roles.GetByName(ctx, requiredRole)
	if err != nil {
		return false, err
	}

	return user.Role.Level >= role.Level, nil
}

func (app *application) checkPostOwnership(requiredRole string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromContext(r)
		post := getPostFromCtx(r)

		if post.UserID == user.ID {
			next.ServeHTTP(w, r)
			return
		}

		allowed, err := app.checkRolePrecedence(r.Context(), user, requiredRole)
		if err != nil {
			app.internalServiceError(w, r, err)
			return
		}

		if !allowed {
			app.forbiddenResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
