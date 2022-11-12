package identity

// identity.Users.GetById(id)

var Users User = User{}
var Roles Role = Role{}

// experimenting with how the api might look to query for users
type Identity struct {
	Users User
	Roles Role
}
