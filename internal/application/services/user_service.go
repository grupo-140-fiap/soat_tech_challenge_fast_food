package services

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/persistance/repositories"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/dto"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) CreateUser(user *dto.CreateUserDTO) error {
	err := u.userRepository.CreateUser(user)

	if err != nil {
		return err
	}
	
	return nil
}