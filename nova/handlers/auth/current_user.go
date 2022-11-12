package auth

import (
	"net/http"

	"github.com/brandonclapp/nova/auth"
	coreHttp "github.com/brandonclapp/nova/http"
)

func CurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		get(w, r)
		return
	}

	http.NotFound(w, r)
}

func get(w http.ResponseWriter, r *http.Request) {
	user, err := auth.GetContextUser(r.Context())

	if err != nil {
		coreHttp.WriteJsonResponse(w, "User not authenticated", http.StatusNotFound)
		return
	}

	coreHttp.WriteJsonResponse(w, user, http.StatusOK)
}
