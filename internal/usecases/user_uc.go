package usecases

import (
	"context"
	"log"

	"github.com/abdullahnettoor/tastybites/internal/auth"
	"github.com/abdullahnettoor/tastybites/internal/models"
	"github.com/abdullahnettoor/tastybites/internal/repo/interfaces"
)

type UserIUsecase interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	LoginUser(ctx context.Context, email, password string) (models.User, error)
	GetUser(ctx context.Context, id int) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, id int) error

	// User operations
	GetOrdersByUser(ctx context.Context, userId int) ([]models.Order, error)

	// Admin operations
	GetOrderByTableId(ctx context.Context, tableId int) (models.Order, error)
}

type UserUsecase struct {
	repo interfaces.Repository
}

func NewUserUsecase(repo interfaces.Repository) UserIUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (u *UserUsecase) LoginUser(ctx context.Context, email, password string) (models.User, error) {
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return models.User{}, err
	}

	log.Printf("User found: %v", user)
	if auth.CompareHashedPassword(user.Password, password) != nil {
		return models.User{}, models.ErrInvalidCredentials
	}

	return user, nil
}

func (u *UserUsecase) CreateUser(ctx context.Context, user models.User) (int, error) {
	var err error
	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	return u.repo.CreateUser(ctx, user)
}

func (u *UserUsecase) GetAvailableTables(ctx context.Context) ([]models.Table, error) {
	return u.repo.GetTablesByStatus(ctx, models.TableStatusAvailable)
}

func (u *UserUsecase) GetUser(ctx context.Context, id int) (models.User, error) {
	return u.repo.GetUserById(ctx, id)
}

func (u *UserUsecase) GetOrdersByUser(ctx context.Context, userId int) ([]models.Order, error) {
	orders, err := u.repo.GetOrdersByUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, models.ErrIsEmpty
	}
	return orders, nil
}

func (u *UserUsecase) UpdateUser(ctx context.Context, user models.User) error {
	return u.repo.UpdateUser(ctx, user)
}

func (u *UserUsecase) DeleteUser(ctx context.Context, id int) error {
	return u.repo.DeleteUser(ctx, id)
}
