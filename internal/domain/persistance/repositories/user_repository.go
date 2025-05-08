package repositories

import (
    "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/dto"
)

type UserRepository interface {
    CreateUser(user *dto.CreateUserDTO) error
}