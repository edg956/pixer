package middleware

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, ok := r.Header["Authorization"]
		if !ok || len(auth) == 0 {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		parts := strings.Split(auth[0], " ")

		if len(parts) != 2 {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		authType, token := parts[0], parts[1]

		validate, err := getValidatorFor(authType)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		userID, err := validate(token)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		var ctx = context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getValidatorFor(tokenType string) (func(token string) (string, error), error) {
	switch tokenType {
	case "dummy":
		return dummyValidate, nil
	default:
		return nil, errors.Errorf("Invalid authorization type: %s", tokenType)
	}
}

func dummyValidate(token string) (string, error) {
	if token == "dummy-token" {
		return "dummy-user-id", nil
	}
	return "", errors.Errorf("Invalid authorization token: %s", token)
}
