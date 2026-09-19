package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pinebit/go-boilerplate/logger"
)

func TestRoutes(t *testing.T) {
	t.Parallel()
	router := NewRouter(logger.NewNoopLogger())
	for _, tc := range []struct {
		method, path string
		status       int
		body         string
	}{
		{"GET", "/", 200, "Hello world!"},
		{"GET", "/missing", 404, "404 page not found"},
		{"POST", "/", 405, "Method Not Allowed"},
		{"GET", "/metrics", 200, "hello_world 1"},
	} {
		response := httptest.NewRecorder()
		router.ServeHTTP(response, httptest.NewRequest(tc.method, tc.path, nil))
		if response.Code != tc.status || !strings.Contains(response.Body.String(), tc.body) {
			t.Fatalf("%s %s: %d %s", tc.method, tc.path, response.Code, response.Body.String())
		}
	}
	other := httptest.NewRecorder()
	NewRouter(logger.NewNoopLogger()).ServeHTTP(other, httptest.NewRequest(http.MethodGet, "/metrics", nil))
	if !strings.Contains(other.Body.String(), "hello_world 0") {
		t.Fatal("router metrics leaked between instances")
	}
}
