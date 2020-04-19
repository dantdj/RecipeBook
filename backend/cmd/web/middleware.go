package main

import (
	"net/http"
)

// Adds some basic security-related headers to a given request.
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("X-XSS-Protection", "1; mode=block")
		writer.Header().Set("X-Frame-Options", "deny")
		writer.Header().Set("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(writer, request)
	})
}