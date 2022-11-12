package auth

// import (
// 	"context"

// 	"github.com/brandonclapp/nova/identity"
// 	"github.com/brandonclapp/nova/graph/model"
// )

// func UserRoles(ctx context.Context, obj *model.User) ([]*model.RoleMembership, error) {
// 	roleMemberships, err := identity.Users.GetRoles(obj.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	modelRoleMemberships := make([]*model.RoleMembership, len(roleMemberships))
// 	for i, roleMembership := range roleMemberships {
// 		modelRoleMemberships[i] = &model.RoleMembership{
// 			RoleID:      roleMembership.RoleID,
// 			DisplayName: roleMembership.DisplayName,
// 			Context:     &roleMembership.Context,
// 			EntityID:    &roleMembership.EntityID,
// 		}
// 	}

// 	return modelRoleMemberships, nil
// }
