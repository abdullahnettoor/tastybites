package usecases

import (
	"context"
	"fmt"

	"github.com/abdullahnettoor/tastybites/internal/models"
	"github.com/abdullahnettoor/tastybites/internal/repo/interfaces"
)

type TableIUsecase interface {
	GetAllTables(ctx context.Context) ([]models.Table, error)
	GetAvailableTables(ctx context.Context) ([]models.Table, error)
	IsTableAvailable(ctx context.Context, tableID int) (bool, error)
	ResetTableStatus(ctx context.Context, tableID int) error
}

type TableUsecase struct {
	repo interfaces.Repository
}

func NewTableUsecase(repo interfaces.Repository) TableIUsecase {
	return &TableUsecase{
		repo: repo,
	}
}

func (m *TableUsecase) GetAvailableTables(ctx context.Context) ([]models.Table, error) {
	tables, err := m.repo.GetTablesByStatus(ctx, models.TableStatusAvailable)
	if err != nil {
		return nil, err
	}
	if len(tables) == 0 {
		return nil, models.ErrIsEmpty
	}
	return tables, nil
}

func (m *TableUsecase) GetAllTables(ctx context.Context) ([]models.Table, error) {
	tables, err := m.repo.GetAllTables(ctx)
	if err != nil {
		return nil, err
	}
	if len(tables) == 0 {
		return nil, models.ErrIsEmpty
	}
	return tables, nil
}

func (m *TableUsecase) IsTableAvailable(ctx context.Context, tableID int) (bool, error) {
	table, err := m.repo.GetTableById(ctx, tableID)
	if err != nil {
		return false, err
	}
	return table.Status == models.TableStatusAvailable, nil
}

func (m *TableUsecase) ResetTableStatus(ctx context.Context, tableID int) error {
	_, err := m.repo.GetTableById(ctx, tableID)
	if err != nil {
		return fmt.Errorf("table not found: %w", err)
	}

	err = m.repo.UpdateOrderStatusByTableId(ctx, tableID, models.OrderStatusCompleted)
	if err != nil {
		return fmt.Errorf("failed to complete orders for table: %w", err)
	}

	err = m.repo.ResetTableToAvailable(ctx, tableID)
	if err != nil {
		return fmt.Errorf("failed to reset table status: %w", err)
	}

	return nil
}
