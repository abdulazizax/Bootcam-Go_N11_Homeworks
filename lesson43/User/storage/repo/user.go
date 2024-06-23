package repo

import (
	m "user/models"
	"context"
)

type UsersStorageI interface {
	CreateUser(ctx context.Context, user m.UserRequest) (m.UserResponse, error)
	GetUserById(ctx context.Context, id string) (m.UserResponse, error)
	GetUsers(ctx context.Context) ([]m.UserResponse, error)
	UpdateUserById(ctx context.Context, id string, user m.UserRequest) (m.UserResponse, error)
	DeleteUserByID(ctx context.Context, id string) error
	GetUserByName(ctx context.Context, name string) (m.UserResponse, error)
}
