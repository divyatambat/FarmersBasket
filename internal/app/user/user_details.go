package user

import (
	"github.com/divyatambat/FarmersBasket/internal/pkg/dto"
	"github.com/divyatambat/FarmersBasket/internal/repository"
)

// Define your user model and related types here (UserType, Address, etc.)

func MapRepoObjectToDto(repoObj repository.User) dto.User {
	// Implement mapping logic based on your user model and DTO definitions
	return dto.User{}
}

func MapDtoObjectToRepo(user dto.User) repository.User {
	// Implement mapping logic based on your user model and DTO definitions
	return repository.User{}
}
