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
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	PhoneNumber int64  `db:"phone_number"`
	UserType    string `db:"user_type"`
}
