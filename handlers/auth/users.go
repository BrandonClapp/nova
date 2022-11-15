package auth

import (
	"net/http"

	coreHttp "github.com/brandonclapp/nova/http"
	"github.com/brandonclapp/nova/identity"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := identity.Users.All()
	coreHttp.WriteJsonResponse(w, &users, http.StatusOK)
}
