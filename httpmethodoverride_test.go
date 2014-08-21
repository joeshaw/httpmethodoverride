package httpmethodoverride

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPMethodOverride(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Method: %s", r.Method)))
	})

	// A call against the regular handler
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost/foo?bar=baz", nil)
	handler.ServeHTTP(rec, req)

	if rec.Body.String() != "Method: GET" {
		t.Fatalf("Expected 'Method: GET', got '%s'", rec.Body.String())
	}

	// Wrapped in the handler, but not overridden
	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "http://localhost/foo?bar=baz", nil)
	Handler(handler).ServeHTTP(rec, req)
	if rec.Body.String() != "Method: GET" {
		t.Fatalf("Expected 'Method: GET', got '%s'", rec.Body.String())
	}

	// Wrapped in the handler and overridden
	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "http://localhost/foo?bar=baz&_method=POST", nil)
	Handler(handler).ServeHTTP(rec, req)
	if rec.Body.String() != "Method: POST" {
		t.Fatalf("Expected 'Method: POST', got '%s'", rec.Body.String())
	}
}
