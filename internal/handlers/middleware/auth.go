package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yohannesgossaye/pkgs/message/localization"
)

type contextKey string

const (
	userIDContextKey contextKey = "user_id"
	emailContextKey  contextKey = "email"
)

// JWTAuthMiddleware validates Bearer tokens and injects user claims into request context.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			localization.SendUnauthorizedResponse(w, "")
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		if tokenString == "" {
			localization.SendUnauthorizedResponse(w, "")
			return
		}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "defaultsecret"
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			localization.SendUnauthorizedResponse(w, "")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			localization.SendUnauthorizedResponse(w, "")
			return
		}

		ctx := r.Context()
		if v, ok := claims["user_id"]; ok {
			ctx = context.WithValue(ctx, userIDContextKey, v)
		}
		if v, ok := claims["email"]; ok {
			ctx = context.WithValue(ctx, emailContextKey, v)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserIDFromContext extracts the user_id claim from context if present.
func GetUserIDFromContext(ctx context.Context) (any, bool) {
	v := ctx.Value(userIDContextKey)
	if v == nil {
		return nil, false
	}
	return v, true
}

// GetEmailFromContext extracts the email claim from context if present.
func GetEmailFromContext(ctx context.Context) (any, bool) {
	v := ctx.Value(emailContextKey)
	if v == nil {
		return nil, false
	}
	return v, true
}
