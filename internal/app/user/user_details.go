package user

import (
	"github.com/divyatambat/FarmersBasket/internal/pkg/dto"
	"github.com/divyatambat/FarmersBasket/internal/repository"
)

func MapRepoObjectToDto(repoObj repository.User) dto.User {
	return dto.User{
		ID:          int64(repoObj.ID),
		Name:        repoObj.Name,
		Email:       repoObj.Email,
		Password:    repoObj.Password,
		PhoneNumber: repoObj.PhoneNumber,
		UserType:    repoObj.UserType,
	}
}

func MapDtoObjectToRepo(user dto.User) repository.User {
	return repository.User{
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		UserType:    user.UserType,
	}
}
