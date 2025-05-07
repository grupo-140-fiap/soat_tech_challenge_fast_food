package services

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/persistance/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}