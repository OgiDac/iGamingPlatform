package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/OgiDac/iGamingPlatform/domain"
	"github.com/OgiDac/iGamingPlatform/utils"
)

func AdminAuthMiddleware(secret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				authHeader := r.Header.Get("Authorization")
				t := strings.Split(authHeader, " ")
				if len(t) == 2 {
					authToken := t[1]
					authorized, err := utils.IsAuthorized(authToken, secret)
					if err != nil {
						utils.JSON(w, 401, domain.ErrorResponse{Message: err.Error()})
						return
					}
					if authorized {
						userID, err := utils.ExtractIDFromToken(authToken, secret)
						if err != nil {
							utils.JSON(w, 401, domain.ErrorResponse{Message: err.Error()})
							return
						}

						userRole, err := utils.ExtractRoleFromToken(authToken, secret)
						if err != nil {
							utils.JSON(w, 401, domain.ErrorResponse{Message: err.Error()})
							return
						}
						if userRole == "" {
							utils.JSON(w, 401, domain.ErrorResponse{Message: "invalid token"})
							return
						}
						if userRole != domain.Admin {
							utils.JSON(w, 401, domain.ErrorResponse{Message: "invalid permisions"})
							return
						}
						// set user id to context
						ctx := context.WithValue(r.Context(), "user_id", userID)
						r = r.WithContext(ctx)
						next.ServeHTTP(w, r)
						return
					}
					utils.JSON(w, 401, domain.ErrorResponse{Message: "unauthorized"})
					return
				}
				utils.JSON(w, 401, domain.ErrorResponse{Message: "unauthorized"})
				return
			})
	}
}
