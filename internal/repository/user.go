package repository

import (
	"context"
)

// UserStorer - methods for interacting with the user data store.
type UserStorer interface {
	GetUserByID(ctx context.Context, userID uint) (User, error)
	ListUsers(ctx context.Context) ([]User, error)
	CreateUser(ctx context.Context, user User) error
}

type User struct {
	ID          int64
	Name        string
	Email       string
	Password    string
	PhoneNumber int64
	UserType    string
}
