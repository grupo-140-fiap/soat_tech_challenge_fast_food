package mysql

import (
    "database/sql"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
    "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.UserRepository {
    return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(user *dto.CreateUserDTO) error {
    query := "INSERT INTO users (first_name, last_name, cpf, email) VALUES (?, ?, ?, ?)"

    _, err := u.db.Exec(query, user.FirstName, user.LastName, user.CPF, user.Email)

    if err != nil {
        return err
    }

    return nil
}