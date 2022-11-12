package auth

import (
	"fmt"

	"github.com/brandonclapp/nova/data"
	"github.com/brandonclapp/nova/identity"
)

// Authenticate checks if the user's password is correct.
// Returns the authenticated user if successful, otherwise
// returns nil
func Authenticate(email string, password string) (*identity.User, error) {

	user := identity.Users.GetUserByEmail(email)

	if user == nil {
		return nil, fmt.Errorf("incorrect email/password combo")
	}

	correctPassword := comparePasswords(user.Password, []byte(password))

	if correctPassword {
		return user, nil
	}

	return nil, fmt.Errorf("incorrect email/password combo")
}

func RegisterUser(user *identity.User) (createdID uint, err error) {
	// TODO: Validate that user.Email is actually a valid email address.

	// Ensure that email isn't already taken
	if existingEmail := identity.Users.GetUserByEmail(user.Email); existingEmail != nil {
		return 0, fmt.Errorf("email already exists")
	}

	// Hash the user's incoming raw password
	password, err := hashAndSalt([]byte(user.Password))
	if err != nil {
		return 0, err
	}

	user.Password = password
	user.IsConfirmed = false

	if tx := data.DB.Create(&user); tx.Error != nil {
		return 0, tx.Error
	}

	return user.ID, nil
}
