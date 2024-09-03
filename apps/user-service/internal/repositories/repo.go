package repositories

import (
    "context"
    "database/sql"
    "user-service/internal/models"
    "errors"
)

type UserRepository interface {
    CreateUser(ctx context.Context, user *models.User) error
    GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type userRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &userRepository{db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
    query := `INSERT INTO users (email, password, created_at, updated_at)
              VALUES ($1, $2, $3, $4) RETURNING id`
    return r.db.QueryRowContext(
        ctx,
        query,
        user.Email,
        user.Password,
        user.CreatedAt,
        user.UpdatedAt,
        ).Scan(&user.ID)
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
    user := &models.User{}
    query := `SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1`
    err := r.db.QueryRowContext(ctx, query, email).Scan(
        &user.ID,
        &user.Email,
        &user.Password,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    if err == sql.ErrNoRows {
        return nil, errors.New("user not found")
    }
    return user, err
}
