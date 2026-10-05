package clogger

import (
	"net/http"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.status = code
	sr.ResponseWriter.WriteHeader(code)
}

func (sr *statusRecorder) Unwrap() http.ResponseWriter {
	return sr.ResponseWriter
}

func (log *Logger) Start(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{w, http.StatusOK}
		next.ServeHTTP(rec, r)
		log.LogRequest(r.Method, r.URL.Path, rec.status, time.Since(start))
	})
}

func (log *Logger) Handler(h http.Handler) http.Handler {
	return log.Start(h)
}
