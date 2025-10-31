package middleware

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yaufdd/project3/internal/auth"
)

type JWTMW struct {
	publicKey *rsa.PublicKey
	issuer    string
}

func NewJWTMW(pub *rsa.PublicKey, issuer string) *JWTMW { return &JWTMW{pub, issuer} }

func (m *JWTMW) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" || !strings.HasPrefix(h, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
			return
		}

		raw := strings.TrimPrefix(h, "Bearer ")

		var claims auth.AcessClaims
		tok, err := jwt.ParseWithClaims(raw, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected sign method")
			}
			return m.publicKey, nil
		}, jwt.WithIssuer(m.issuer))
		if err != nil || !tok.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}
		c.Set("uid", claims.UserID)
		c.Set("roles", claims.Roles)
		c.Next()
	}
}

func (m *JWTMW) RequiredRoles(need ...string) gin.HandlerFunc {
	set := map[string]struct{}{}
	for _, r := range need {
		set[r] = struct{}{}
	}
	return func(c *gin.Context) {
		v, ok := c.Get("roles")
		if !ok {
			c.AbortWithStatusJSON(403, gin.H{"error": "forbidden"})
			return
		}
		roles, ok := v.([]string)
		if !ok {
			c.AbortWithStatusJSON(403, gin.H{"error": "forbidden"})
			return
		}

		for _, have := range roles {
			if _, ok := set[have]; ok {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
	}
}
