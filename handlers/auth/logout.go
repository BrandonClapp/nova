package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/brandonclapp/nova/auth"
	"github.com/brandonclapp/nova/sessions"
)

func sendJson[T any](w http.ResponseWriter, status int, item T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	b, err := json.Marshal(item)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	io.WriteString(w, string(b))
}

// TODO: Move this out if this is something that I want to standardize on
type ActionResponse struct {
	Success bool
	Message string
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure that logout is a POST
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	context := r.Context()
	user, _ := auth.GetContextUser(context)

	successResponse := ActionResponse{Success: true, Message: "User logged out"}

	if user == nil {
		// If user is already logged out, just return success true
		sendJson(w, 200, successResponse)
		return
	}

	err := sessions.DeleteSession(context, user.ID)

	if err != nil {
		sendJson(w, http.StatusInternalServerError, ActionResponse{
			Success: false,
			Message: "Unable to delete user session",
		})
		return
	}

	sendJson(w, http.StatusOK, successResponse)
}
