package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoot(t *testing.T) {
	rr := httptest.NewRecorder()
	newMux().ServeHTTP(rr, httptest.NewRequest("GET", "/", nil))
	if got := rr.Body.String(); got != "hello from my container\n" {
		t.Errorf("GET / = %q, want %q", got, "hello from my container\n")
	}
}

func TestHealthz(t *testing.T) {
	rr := httptest.NewRecorder()
	newMux().ServeHTTP(rr, httptest.NewRequest("GET", "/healthz", nil))
	if rr.Code != http.StatusOK || rr.Body.String() != "ok\n" {
		t.Errorf("GET /healthz = %d %q, want 200 %q", rr.Code, rr.Body.String(), "ok\n")
	}
}

func TestPortDefault(t *testing.T) {
	t.Setenv("PORT", "")
	if got := port(); got != "8080" {
		t.Errorf("port() = %q, want 8080", got)
	}
}
