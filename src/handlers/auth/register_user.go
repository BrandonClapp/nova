package auth

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/brandonclapp/nova/auth"
// 	"github.com/brandonclapp/nova/identity"
// 	"github.com/brandonclapp/nova/graph/model"
// )

// func RegisterUser(ctx context.Context, input model.NewRegister) (*model.User, error) {
// 	if user, _ := auth.GetContextUser(ctx); user != nil {
// 		// Already logged in
// 		return nil, fmt.Errorf("already authenticated")
// 	}

// 	createdAt := time.Now()
// 	user := &identity.User{
// 		Name:        input.Name,
// 		Email:       input.Email,
// 		Password:    input.Password,
// 		CreatedAt:   &createdAt,
// 		IsActive:    true,
// 		IsConfirmed: false,
// 		LastAuth:    nil,
// 	}

// 	id, err := auth.RegisterUser(user)

// 	if err != nil {
// 		log.Printf("error while registering user")
// 		return nil, fmt.Errorf("error while registering user: %v", err)
// 	}

// 	modelUser := &model.User{
// 		ID:          id,
// 		Name:        user.Name,
// 		Email:       user.Email,
// 		CreatedAt:   user.CreatedAt,
// 		IsActive:    user.IsActive,
// 		IsConfirmed: user.IsConfirmed,
// 		LastAuth:    user.LastAuth,
// 	}

// 	return modelUser, nil
// }
