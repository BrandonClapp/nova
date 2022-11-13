package cors

import (
	"net/http"
	"strings"
)

var allowedHeaders = []string{
	"Origin",
	"X-Requested-With",
	"Content-Type",
	"Accept",
	"Authorization",
	"sentry-trace",
}

// TODO: should be configurable
var allowedOrigins = []string{
	"http://localhost:3000",
	"https://brackets.gg",
}

// access control and  CORS middleware
func AccessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")

		var isAllowed bool = false
		for _, allowed := range allowedOrigins {
			if allowed == origin {
				isAllowed = true
				break
			}
		}

		if isAllowed {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(allowedHeaders[:], ","))
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
