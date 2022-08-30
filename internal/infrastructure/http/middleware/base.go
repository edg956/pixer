package middleware

import (
	"net/http"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

func ApplyMiddleware(f http.HandlerFunc) http.HandlerFunc {
	var middlewares []middleware
	return chainMiddleware(f, middlewares...)
}

func chainMiddleware(f http.HandlerFunc, middle ...middleware) http.HandlerFunc {
	if len(middle) == 0 {
		return f
	}

	return middle[0](chainMiddleware(f, middle[1:cap(middle)]...))
}
