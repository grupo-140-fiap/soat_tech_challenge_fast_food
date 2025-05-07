package repositories

import (
    "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
)

type UserRepository interface {
    CreateUser(user *entities.User) error
}