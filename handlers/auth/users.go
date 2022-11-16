package auth

import (
	"net/http"

	coreHttp "github.com/brandonclapp/nova/http"
	"github.com/brandonclapp/nova/identity"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getUsersHandler(w, r)
		return
	}

	if r.Method == "POST" {
		createNewUser(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := identity.Users.All()
	coreHttp.WriteJsonResponse(w, &users, http.StatusOK)
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		coreHttp.WriteJsonResponse(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user *identity.User
	coreHttp.ParseBody(w, r, &user)

	err := identity.Users.Create(user)

	if err != nil {
		coreHttp.WriteJsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	coreHttp.WriteJsonResponse(w, &user, http.StatusOK)
}
