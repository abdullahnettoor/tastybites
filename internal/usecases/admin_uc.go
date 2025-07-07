package usecases

import (
	"context"

	"github.com/abdullahnettoor/tastybites/internal/models"
)

func (u *UserUsecase) GetOrderByTableId(ctx context.Context, tableId int) (models.Order, error) {
	return u.repo.GetOrderByTableId(ctx, tableId)
}

func (u *UserUsecase) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	return u.repo.GetAllOrders(ctx)
}
