package models

type NewUserRequest struct {
	FirstName     *string `json:"firstName"`
	LastName      *string `json:"lastName"`
	Email         string  `json:"email"`
	Country       *string `json:"country"`
	StreetAddress *string `json:"streetAddress"`
	City          *string `json:"city"`
	State         *string `json:"state"`
	Zip           *string `json:"zip"`
	IsActive      bool    `json:"isActive"`
	IsConfirmed   bool    `json:"isConfirmed"`
	Password      string  `json:"password"`
}
