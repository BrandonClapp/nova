package identity

import "github.com/brandonclapp/nova/data"

func AutoMigrate() {
	// create tables for role, user, and join table user_roles
	data.DB.AutoMigrate(&Role{}, &User{})

	var role Role
	result := data.DB.Where(&Role{DisplayName: "user"}).First(&role)
	if result.RowsAffected == 0 {
		role = Role{DisplayName: "user"}
		data.DB.Create(&role)
	}

	var user User
	result = data.DB.First(&user, "Email = ?", "hey@brandonclapp.com")
	if result.RowsAffected == 0 {
		// hashed password = abcdefg
		name := "Brandon"
		user := &User{
			Name:        &name,
			Email:       "hey@brandonclapp.com",
			Password:    "$2a$04$htjPGT7Yxzv1xU.EQfCBmOvlVry/wxUh3Q0bCHYPZbcWTTn5X47zu",
			IsActive:    true,
			IsConfirmed: true,
			LastAuth:    nil,
			Roles:       []Role{role},
		}

		if tx := data.DB.Create(&user); tx.Error != nil {
			panic(tx.Error)
		}
	}
}
