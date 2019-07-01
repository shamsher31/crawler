
package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPopulateHeaders(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	})

	rr := httptest.NewRecorder()

	// func Headers(h http.Handler) http.Handler
	// Returns content type headers
	handler := Headers(testHandler)
	handler.ServeHTTP(rr, req)

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("content type header does not match: got %v want %v", ctype, "application/json")
	}
}