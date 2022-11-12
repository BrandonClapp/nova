package identity

import "github.com/brandonclapp/nova/data"

type Role struct {
	data.GormModel
	DisplayName string `json:"displayName"`
	Context     string `json:"context"`
}

func (role *Role) All() ([]*Role, error) {
	var roles []*Role

	tx := data.DB.Find(&role)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return roles, nil
}
