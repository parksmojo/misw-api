package middleware

import (
	"net/http"
)

func ApplyTo(handler http.HandlerFunc) http.Handler {
	return http.HandlerFunc(
		LogRequest(handler),
	)
}