package auth

import (
	"net/http"

	"github.com/brandonclapp/nova/auth"
	"github.com/brandonclapp/nova/handlers/auth/models"
	coreHttp "github.com/brandonclapp/nova/http"
	"github.com/brandonclapp/nova/sessions"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		post(w, r)
		return
	}

	http.NotFound(w, r)
}

func post(w http.ResponseWriter, r *http.Request) {

	var input = &models.LoginRequest{}
	if err := coreHttp.ParseBody(w, r, input); err != nil {
		// ParseBody writes error headers + body
		return
	}

	user, err := auth.Authenticate(input.Email, input.Password)

	if user != nil {
		// sessions.CreateSession will write session cookie to response header
		err = sessions.CreateSession(r.Context(), user.ID)

		if err != nil {
			coreHttp.WriteJsonResponse(w, "Error creating session", http.StatusInternalServerError)
			return
		}

		coreHttp.WriteJsonResponse(w, user, http.StatusOK)
		return
	}

	// General authentication error. Probably incorrect email/password.
	coreHttp.WriteJsonResponse(w, "Incorrect email address and/or password.", http.StatusBadRequest)
}
