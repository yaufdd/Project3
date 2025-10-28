package auth

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthJWT struct {
	privateKey *rsa.PrivateKey
	issuer     string
	accessTTL  time.Duration
}

func NewAuthJWT(priv *rsa.PrivateKey, issuer string, ttl time.Duration) *AuthJWT {
	return &AuthJWT{privateKey: priv, issuer: issuer, accessTTL: ttl}
}

type Claims struct {
	UserID int64 `json:"uid"`
	Roles  []string
	jwt.RegisteredClaims
}

func (s *AuthJWT) IssueAccessToken(userID int64, roles []string) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID: userID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer,
			Subject:   fmt.Sprint(userID),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.accessTTL)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(s.privateKey)
}
