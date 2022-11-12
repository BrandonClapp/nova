package identity

import (
	"time"

	"github.com/brandonclapp/nova/data"
)

// User is the identity package representation of a user.
type User struct {
	data.GormModel
	Name        *string    `json:"name"`
	Email       string     `json:"email"`
	Password    string     `json:"-"`
	IsActive    bool       `json:"isActive"`
	IsConfirmed bool       `json:"isConfirmed"`
	LastAuth    *time.Time `json:"lastAuth"`
	CreatedAt   time.Time  `json:"createdAt"`
	Roles       []Role     `json:"roles" gorm:"many2many:user_roles;"`
}

func (user *User) All() ([]*User, error) {
	var users []*User

	tx := data.DB.Find(&users)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

// GetUserByID gets the user from the database by ID
func (user *User) GetUserByID(userId uint) *User {
	var u *User
	tx := data.DB.Preload("Roles").First(&u, userId)

	if tx.Error != nil {
		return nil
	}

	return u
}

// GetUserByEmail gets the user from the database by email address.
func (user *User) GetUserByEmail(email string) *User {
	var u *User
	tx := data.DB.Preload("Roles").Where("Email = ?", email).First(&u)

	if tx.Error != nil {
		return nil
	}

	return u
}
