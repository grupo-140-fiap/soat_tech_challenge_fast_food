package mysql

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/persistance/repositories"
    "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
)

type UserRepository struct {}

func NewUserRepository() repositories.UserRepository {
    return &UserRepository{}
}

func (u *UserRepository) CreateUser(user *entities.User) error {

    return nil
}