package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const testSecret = "test"

type TestAuthenticator struct{}

func (a *TestAuthenticator) GenerateToken(jwt.Claims) (string, error) {
	claims := jwt.MapClaims{
		"aud": "test",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iss": "test",
		"sub": "1",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(testSecret))
}

func (a *TestAuthenticator) ValidateToken(token string) (*jwt.Token, error) {
	keyFunc := func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(testSecret), nil
	}

	return jwt.Parse(
		token,
		keyFunc,
		jwt.WithIssuer("test"),
		jwt.WithAudience("test"),
		jwt.WithExpirationRequired(),
	)
}
