package usecases

import (
	"context"

	"github.com/abdullahnettoor/tastybites/internal/models"
	"github.com/abdullahnettoor/tastybites/internal/repo/interfaces"
)

type TableIUsecase interface {
	GetAllTables(ctx context.Context) ([]models.Table, error)
	GetAvailableTables(ctx context.Context) ([]models.Table, error)
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
