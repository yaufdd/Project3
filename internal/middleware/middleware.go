package middleware

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
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

func (m *JWTMW) RequiredAccountOnwer() gin.HandlerFunc {
	return func(c *gin.Context) {
		authUserID, exist := c.Get("uid")
		if !exist {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthrized"})
			fmt.Println("1")
			return
		}
		fmt.Println(reflect.TypeOf(authUserID))
		uidFromToken, ok := authUserID.(int64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthrized"})
			fmt.Println("2")
			return
		}
		bodyBytes, err := io.ReadAll(c.Request.Body)
		fmt.Println(bodyBytes)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthrized"})
			fmt.Println("3")
			return
		}

		var payload map[string]any
		if len(bodyBytes) > 0 {
			if err := json.Unmarshal(bodyBytes, &payload); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
				fmt.Println("4")
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "empty body"})
			fmt.Println("5")
			return
		}
		rawBodyUID, ok := payload["user_id"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user_id missing in body"})
			fmt.Println("6")
			return
		}

		var uidFromBody int64
		switch v := rawBodyUID.(type) {
		case float64:
			uidFromBody = int64(v)
		case int64:
			uidFromBody = v
		default:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user_id has wrong type"})
			fmt.Println("7")
			return
		}
		if uidFromBody != uidFromToken {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			fmt.Println("8")
			return
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
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
