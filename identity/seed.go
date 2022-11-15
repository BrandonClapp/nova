package identity

import "github.com/brandonclapp/nova/data"

func AutoMigrate() {
	// create tables for role, user, and join table user_roles
	data.DB.AutoMigrate(&Role{}, &User{})

	var userRole Role
	result := data.DB.Where(&Role{DisplayName: "user"}).First(&userRole)
	if result.RowsAffected == 0 {
		userRole = Role{DisplayName: "user"}
		data.DB.Create(&userRole)
	}

	var systemAdminRole Role
	result = data.DB.Where(&Role{DisplayName: "system admin"}).First(&systemAdminRole)
	if result.RowsAffected == 0 {
		systemAdminRole = Role{DisplayName: "system admin"}
		data.DB.Create(&systemAdminRole)
	}

	var user1 User
	result = data.DB.First(&user1, "Email = ?", "hey@brandonclapp.com")
	if result.RowsAffected == 0 {
		// hashed password = abcdefg
		name := "Brandon"
		user1 := &User{
			Name:        &name,
			Email:       "hey@brandonclapp.com",
			Password:    "$2a$04$htjPGT7Yxzv1xU.EQfCBmOvlVry/wxUh3Q0bCHYPZbcWTTn5X47zu",
			IsActive:    true,
			IsConfirmed: true,
			LastAuth:    nil,
			Roles:       []Role{systemAdminRole},
		}

		if tx := data.DB.Create(&user1); tx.Error != nil {
			panic(tx.Error)
		}
	}

	var user2 User
	result = data.DB.First(&user2, "Email = ?", "user@brandonclapp.com")
	if result.RowsAffected == 0 {
		// hashed password = abcdefg
		name := "Brandon (Normal)"
		user2 := &User{
			Name:        &name,
			Email:       "user@brandonclapp.com",
			Password:    "$2a$04$htjPGT7Yxzv1xU.EQfCBmOvlVry/wxUh3Q0bCHYPZbcWTTn5X47zu",
			IsActive:    true,
			IsConfirmed: true,
			LastAuth:    nil,
			Roles:       []Role{userRole},
		}

		if tx := data.DB.Create(&user2); tx.Error != nil {
			panic(tx.Error)
		}
	}
}
