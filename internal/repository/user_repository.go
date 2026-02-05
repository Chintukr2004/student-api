package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Chintukr2004/student-api/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}
func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return r.DB.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.PasswordHash,
	).Scan(&user.ID, &user.CreatedAt)
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, name, email, role, password_hash, created_at
		FROM users
		WHERE email = $1
	`

	var user models.User

	err := r.DB.QueryRowContext(ctx, query, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
