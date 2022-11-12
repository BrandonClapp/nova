package auth

// import (
// 	"context"

// 	"github.com/brandonclapp/nova/identity"
// 	"github.com/brandonclapp/nova/graph/model"
// )

// func Users(ctx context.Context) ([]*model.User, error) {
// 	users, err := identity.Users.All()
// 	if err != nil {
// 		return nil, err
// 	}

// 	modelUsers := make([]*model.User, len(users))
// 	for i, user := range users {
// 		modelUsers[i] = &model.User{
// 			ID:          user.ID,
// 			Name:        user.Name,
// 			Email:       user.Email,
// 			IsActive:    user.IsActive,
// 			IsConfirmed: user.IsConfirmed,
// 			CreatedAt:   user.CreatedAt,
// 			LastAuth:    user.LastAuth,
// 		}
// 	}

// 	return modelUsers, nil
// }
