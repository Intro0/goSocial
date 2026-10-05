package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAuthenticator struct {
	secret   []byte
	issuer   string
	audience string
}

func NewJWTAuthenticator(secret, issuer, audience string) *JWTAuthenticator {
	return &JWTAuthenticator{
		secret:   []byte(secret),
		issuer:   issuer,
		audience: audience,
	}
}

func (a *JWTAuthenticator) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.secret)
}

func (a *JWTAuthenticator) ValidateToken(token string) (*jwt.Token, error) {
	keyFunc := func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return a.secret, nil
	}

	options := []jwt.ParserOption{
		jwt.WithIssuer(a.issuer),
		jwt.WithAudience(a.audience),
		jwt.WithExpirationRequired(),
	}

	return jwt.Parse(token, keyFunc, options...)
}
