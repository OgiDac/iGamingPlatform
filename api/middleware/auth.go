package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/OgiDac/iGamingPlatform/utils"
)

func JwtAuthMiddleware(secret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				authHeader := r.Header.Get("Authorization")
				t := strings.Split(authHeader, " ")
				if len(t) == 2 {
					authToken := t[1]
					authorized, err := utils.IsAuthorized(authToken, secret)
					if err != nil {
						utils.JSON(w, 401, errors.New(err.Error()))
						return
					}
					if authorized {
						userID, err := utils.ExtractIDFromToken(authToken, secret)
						if err != nil {
							utils.JSON(w, 401, errors.New(err.Error()))
							return
						}
						// set user id to context
						ctx := context.WithValue(r.Context(), "user_id", userID)
						r = r.WithContext(ctx)
						next.ServeHTTP(w, r)
						return
					}
					utils.JSON(w, 401, errors.New("unauthorized"))
					return
				}
				utils.JSON(w, 401, errors.New("unauthorized"))
				return
			})
	}
}
