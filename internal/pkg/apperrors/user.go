package apperrors

import (
	"errors"
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

// UserAlreadyExistsError indicates an attempt to create a duplicate user.
type UserAlreadyExistsError struct {
	Email string
}

func (e UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("User with email %s already exists", e.Email)
}

func NewValidator() *validator.Validate {
	// Create and configure a validator instance
	validate := validator.New()
	// ... add validation rules and configurations
	return validate
}

// ErrInvalidCredentials represents invalid login credentials.
var ErrInvalidCredentials = errors.New("invalid email or password")

// ErrInvalidUserID indicates an invalid user ID.
var ErrInvalidUserID = errors.New("invalid user ID")

var WrapError = errors.New("Error while generating token")
