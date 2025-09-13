package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type authContextKeyType struct{}

type authContextValue struct {
	UserID int64
}

func (ct authContextKeyType) Set(ctx context.Context, value authContextValue) context.Context {
  new_ctx := context.WithValue(ctx, ct, value)
  return new_ctx
}

func (ct *authContextKeyType) Get(ctx context.Context) (authContextValue, error) {
  value, ok := ctx.Value(*ct).(authContextValue)
  if !ok {
    return authContextValue{}, errors.New("Cannot get context value")
  }
  return value, nil
}

var authContextKey authContextKeyType

func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			var err error

			cookie, err := r.Cookie("social-auth-token")

			if err != nil {
				AuthError(w)
				return
			}

			parts := strings.Split(cookie.Value, " ")

			if len(parts) != 2 || parts[0] != "Bearer" {
				AuthError(w)
				return
			}

			//do some verification with the parts[1] aka the jwt string token
			token, err := jwt.Parse(
				parts[1],
				func(token *jwt.Token) (any, error) {
					return app.config.auth.jwtSecret, nil
				},
				jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
			)

			if err != nil {
				AuthError(w)
				return
			}

			sub, err := token.Claims.GetSubject()

			if err != nil {
				AuthError(w)
				return 
			}

			userId, err := strconv.Atoi(sub)

			if err != nil {
				AuthError(w)
				return 
			}

			ctx := authContextKey.Set(
				r.Context(), 
				authContextValue{
					UserID: int64(userId),
				})

				next.ServeHTTP(w, r.WithContext(ctx))

			})
		}
