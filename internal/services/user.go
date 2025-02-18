package services

import (
	"api-dev/domain"
	"api-dev/dto"
	"context"
)

type userService struct {
	userRepo domain.UserRepository
}

func NewUser(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepo: userRepository,
	}
}

// Index implements domain.UserService.
func (u *userService) Index(ctx context.Context) ([]dto.UserData, error) {
	users, err := u.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var userData []dto.UserData
	for _, v := range users {
		userData = append(userData, dto.UserData{
			Id:    v.Id,
			Name:  v.Name,
			Email: v.Email,
		})
	}
	return userData, nil
}
