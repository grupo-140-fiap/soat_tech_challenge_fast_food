package repositories

import (
    "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
)

type UserRepository interface {
    CreateUser(user *dto.CreateUserDTO) error
}