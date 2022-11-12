package auth

// import (
// 	"context"
// 	"fmt"

// 	"github.com/brandonclapp/nova/auth"
// 	"github.com/brandonclapp/nova/identity"
// 	"github.com/brandonclapp/nova/graph/model"
// )

// func Roles(ctx context.Context) ([]*model.Role, error) {
// 	if user, _ := auth.GetContextUser(ctx); user == nil {
// 		return nil, fmt.Errorf("must be authenticated")
// 	}

// 	roles, err := identity.Roles.All()
// 	if err != nil {
// 		return nil, err
// 	}

// 	modelRoles := make([]*model.Role, len(roles))
// 	for i, role := range roles {
// 		modelRoles[i] = &model.Role{
// 			ID:          role.ID,
// 			DisplayName: role.DisplayName,
// 			Context:     &role.Context,
// 		}
// 	}

// 	return modelRoles, nil
// }
