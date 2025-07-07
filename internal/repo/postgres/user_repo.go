package pgrepo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/abdullahnettoor/tastybites/internal/models"
)

// User operations
func (r *repository) CreateUser(ctx context.Context, user models.User) (int, error) {
	// Insert user into the db
	query := `INSERT INTO public.users (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id`
	var userID int
	if user.Role == "" {
		user.Role = string(models.UserRoleUser)
	}
	err := r.DB.QueryRowContext(ctx, query, user.Name, user.Email, user.Password, user.Role).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	return userID, nil
}

func (r *repository) GetUserById(ctx context.Context, id int) (models.User, error) {
	query := `SELECT id, name, email, role, created_at, updated_at FROM public.users WHERE id = $1`
	var user models.User
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found with id %d: %w", id, err)
		}
		return models.User{}, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	query := `SELECT id, name, email, password, role, created_at, updated_at FROM public.users WHERE email = $1`
	var user models.User
	err := r.DB.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found with email %s: %w", email, err)
		}
		return models.User{}, fmt.Errorf("failed to get user by email: %w", err)
	}
	return user, nil
}

func (r *repository) UpdateUser(ctx context.Context, user models.User) error {
	return fmt.Errorf("not implemented")
}

func (r *repository) DeleteUser(ctx context.Context, id int) error {
	return fmt.Errorf("not implemented")
}

// Admin operations
func (r *repository) GetAdminById(ctx context.Context, id int) (models.User, error) {
	return models.User{}, fmt.Errorf("not implemented")
}
