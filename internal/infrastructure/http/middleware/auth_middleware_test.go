package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddleware(t *testing.T) {
	var called = false

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		val := r.Context().Value("userID")
		if val == nil {
			t.Error("userID not present")
		}
		valStr, ok := val.(string)
		if !ok {
			t.Error("not string")
		}
		if valStr != "dummy-user-id" {
			t.Error("wrong userID")
		}
	})

	handlerToTest := AuthMiddleware(nextHandler)

	// create a mock request to use
	req := httptest.NewRequest("GET", "http://www.test.com", nil)
	req.Header["Authorization"] = []string{
		"dummy dummy-token",
	}

	// call the handler using a mock response recorder (we'll not use that anyway)
	handlerToTest.ServeHTTP(httptest.NewRecorder(), req)

	if !called {
		t.Errorf("Mock middelware never called")
	}
}

func TestAuthMiddlewareReturnsUnauthorizedIfNoHeader(t *testing.T) {
	// create a mock request to use
	req := httptest.NewRequest("GET", "http://www.test.com", nil)
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})
	handlerToTest := AuthMiddleware(handler)
	// call the handler using a mock response recorder (we'll not use that anyway)
	handlerToTest.ServeHTTP(rec, req)

	res := rec.Result()

	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected error code %d. Got %d instead.", http.StatusUnauthorized, res.StatusCode)
	}
}
