/*
	Auth has a direct dependency on gorilla sessions to retreive session information
*/

package auth

import (
	"context"
	"fmt"

	"github.com/brandonclapp/nova/identity"
	"github.com/brandonclapp/nova/sessions"
)

// GetContextUser gets the *User from the context.
// If the session cookie has a userId value, the `InjectUserMiddleware` will
// attach the user to the request context, making it accessible from here.
func GetContextUser(ctx context.Context) (*identity.User, error) {
	user := ctx.Value(sessions.UserCtxKey)

	if user != nil {
		// Ensure that the user is valid in the DB
		u := user.(*identity.User)
		foundUser := identity.Users.GetUserByID(u.ID)
		user = foundUser
		return u, nil
	}

	return nil, fmt.Errorf("no user on context")
}
