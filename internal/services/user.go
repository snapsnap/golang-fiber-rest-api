package services

import (
	"api-dev/domain"
	"api-dev/dto"
	"context"
	"database/sql"
	"time"
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
func (u *userService) Index(ctx context.Context, limit int, page int) ([]dto.UserData, error) {
	users, err := u.userRepo.FindAll(ctx, limit, page)
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

// Create implements domain.UserService.
func (u *userService) Create(ctx context.Context, req dto.RegisterUserRequest) error {
	user := domain.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}
	return u.userRepo.Save(ctx, &user)
}
