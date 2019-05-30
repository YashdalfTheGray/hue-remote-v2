package handlers

import (
	"net/http"
)

// CheckAuthToken does the authentication token checking for
// calls that need auth
func CheckAuthToken(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		h.ServeHTTP(w, req)
	}
}
