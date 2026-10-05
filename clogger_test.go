package clogger

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLog(t *testing.T) {
	logger := New(nil)
	ping := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})
	handler := logger.Start(ping)
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}
	if body := rr.Body.String(); body != "pong" {
		t.Errorf("expected body \"pong\", got %q", body)
	}
}
