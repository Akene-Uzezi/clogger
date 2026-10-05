package clogger

import (
	"net/http"
	"testing"
)

func TestMain(m *testing.M) {
	logger := New(nil)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", logger.Start(ping))
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ping"))
}
