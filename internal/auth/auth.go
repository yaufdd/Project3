package auth

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthJWT struct {
	privateKey *rsa.PrivateKey
	issuer     string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewAuthJWT(priv *rsa.PrivateKey, issuer string, ttl time.Duration) *AuthJWT {
	return &AuthJWT{privateKey: priv, issuer: issuer, accessTTL: ttl}
}

type AcessClaims struct {
	UserID int64 `json:"uid"`
	Roles  []string
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	UID  int64 `json:"uid"`
	Type string
	JTI  string
	jwt.RegisteredClaims
}

func (a *AuthJWT) IssueAccessToken(userID int64, roles []string) (string, time.Time, error) {
	now := time.Now()
	exp := now.Add(a.accessTTL)
	claims := AcessClaims{
		UserID: userID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    a.issuer,
			Subject:   fmt.Sprint(userID),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signed, err := token.SignedString(a.privateKey)
	return signed, exp, err
}

func (a *AuthJWT) IssueRefreshToken(uid int64) (token string, jti string, exp time.Time, err error) {
	now := time.Now()
	exp = now.Add(a.refreshTTL)
	jti = uuid.NewString()

	claims := RefreshClaims{
		UID:  uid,
		Type: "refresh",
		JTI:  jti,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    a.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signed, err := t.SignedString(a.privateKey)
	return signed, jti, exp, err
}
