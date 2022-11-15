package auth

import (
	"net/http"

	coreHttp "github.com/brandonclapp/nova/http"
)

// AsSystemAdmin is a middleware that ensure that the context user is a system admin
// before executing the inner handler.
func AsSystemAdmin(handler func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Requires admin
		user, err := GetContextUser(r.Context())

		// If user is unauthenticated
		if err != nil {
			coreHttp.WriteJsonResponse(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If user is not system admin
		if !IsSystemAdmin(user) {
			coreHttp.WriteJsonResponse(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		handler(w, r)
	})
}
