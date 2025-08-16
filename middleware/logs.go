package middleware

import (
	"log"
	"net/http"
	"time"
)

// LogRequest logs important details about each incoming HTTP request.
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("Started %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("Completed %s in %v", r.URL.Path, duration)
	})
}
