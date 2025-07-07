package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/abdullahnettoor/tastybites/internal/auth"
	"github.com/abdullahnettoor/tastybites/internal/utils"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		tokenParts := strings.Split(tokenStr, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" || tokenParts[1] == "" {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, "invalid token format")
			return
		}
		secretKey := os.Getenv("JWT_SECRET_KEY")
		isValid, claims := auth.IsValidToken(secretKey, tokenParts[1])
		if !isValid {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, "invalid token")
			return
		}
		customClaims, ok := claims.(*auth.CustomClaims)
		if !ok {
			utils.WriteErrorResponse(w, http.StatusForbidden, "invalid token claims type")
			return
		}
		role := customClaims.Role
		userId := customClaims.UserId
		if userId == 0 {
			utils.WriteErrorResponse(w, http.StatusForbidden, "userId not found in token claims")
			return
		}
		fmt.Println("Claims:", customClaims)
		newCtx := context.WithValue(r.Context(), "userId", userId)
		newCtx = context.WithValue(newCtx, "role", role)
		r = r.WithContext(newCtx)

		next.ServeHTTP(w, r)
	})
}
