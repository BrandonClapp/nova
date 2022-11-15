package auth

import (
	"net/http"

	"github.com/brandonclapp/nova/auth"
	coreHttp "github.com/brandonclapp/nova/http"
	"github.com/brandonclapp/nova/identity"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Requires admin
	user, err := auth.GetContextUser(r.Context())

	// If user is unauthenticated
	if err != nil {
		coreHttp.WriteJsonResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// If user is not system admin
	if !auth.IsSystemAdmin(user) {
		coreHttp.WriteJsonResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	users, _ := identity.Users.All()
	coreHttp.WriteJsonResponse(w, &users, http.StatusOK)
}
