package usecases

import (
	"context"

	"github.com/abdullahnettoor/tastybites/internal/models"
	"github.com/abdullahnettoor/tastybites/internal/repo/interfaces"
)

type OrderIUsecase interface {
	CreateOrder(ctx context.Context, order models.Order) (int, error)
	GetOrderById(ctx context.Context, id int) (models.Order, error)
	GetAllOrders(ctx context.Context) ([]models.Order, error)
	GetOrdersByUser(ctx context.Context, userId int) ([]models.Order, error)
	UpdateOrder(ctx context.Context, order models.Order) error
	DeleteOrder(ctx context.Context, id int) error
}

type OrderUsecase struct {
	repo interfaces.Repository
}

func NewOrderUsecase(repo interfaces.Repository) OrderIUsecase {
	return &OrderUsecase{
		repo: repo,
	}
}

func (o *OrderUsecase) CreateOrder(ctx context.Context, order models.Order) (int, error) {
	order.CalculateTotalPrice() // Calculate total price before saving
	return o.repo.CreateOrder(ctx, order)
}

func (o *OrderUsecase) GetOrderById(ctx context.Context, id int) (models.Order, error) {
	return o.repo.GetOrderById(ctx, id)
}

func (o *OrderUsecase) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	return o.repo.GetAllOrders(ctx)
}

func (o *OrderUsecase) GetOrdersByUser(ctx context.Context, userId int) ([]models.Order, error) {
	return o.repo.GetOrdersByUser(ctx, userId)
}

func (o *OrderUsecase) UpdateOrder(ctx context.Context, order models.Order) error {
	order.CalculateTotalPrice() // Recalculate total price before updating
	return o.repo.UpdateOrder(ctx, order)
}

func (o *OrderUsecase) DeleteOrder(ctx context.Context, id int) error {
	return o.repo.DeleteOrder(ctx, id)
}
