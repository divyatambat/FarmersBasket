package user

import (
	"context"
	"time"

	"github.com/divyatambat/FarmersBasket/internal/pkg/apperrors"
	"github.com/divyatambat/FarmersBasket/internal/pkg/dto"
	"github.com/divyatambat/FarmersBasket/internal/repository"

	"github.com/golang-jwt/jwt"
)

type service struct {
	userRepo repository.UserStorer
}

type Service interface {
	GetUserByID(ctx context.Context, tx repository.Transaction, userID uint) (dto.User, error)
	ListUsers(ctx context.Context) ([]dto.User, error)
}

func NewService(userRepo repository.UserStorer) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (ps *service) GetUserByID(ctx context.Context, tx repository.Transaction, userID uint) (dto.User, error) {
	var userDB repository.User

	// Map the retrieved user data to the DTO
	user := dto.User{
		ID:          userDB.ID,
		Name:        userDB.Name,
		Email:       userDB.Email,
		PhoneNumber: userDB.PhoneNumber,
		UserType:    userDB.UserType,
	}

	// Generate JWT token (assuming user ID is unique)
	_, err := generateToken(uint(userID), string([]byte("6a:dc:01:f3:cd:95:4a:3f:27:f1:34:af:d2:2d:6a:cb:ef:5b:b4:62:c2:08:87:76:88:ca:08:e1:61:60:f6:13")))
	if err != nil {
		return dto.User{}, apperrors.WrapError
	}
	// Return the user with the JWT token
	return user, nil
}

func generateToken(userID uint, signingKey string) (string, error) {
	// Set up claims for the token
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// signing key
	// signingKey = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	// Generate the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (*service) ListUsers(ctx context.Context) ([]dto.User, error) {
	panic("unimplemented")
}
