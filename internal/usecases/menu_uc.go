package usecases

import (
	"context"

	"github.com/abdullahnettoor/tastybites/internal/models"
	"github.com/abdullahnettoor/tastybites/internal/repo/interfaces"
)

type MenuIUsecase interface {
	GetAllMenuItems(ctx context.Context) ([]models.MenuItem, error)
	// CreateMenuItem(ctx context.Context, item models.MenuItem) (int, error)
	// GetMenuItemById(ctx context.Context, id int) (models.MenuItem, error)
	// GetMenuItemsByCategory(ctx context.Context, category string) ([]models.MenuItem, error)
	// UpdateMenuItem(ctx context.Context, item models.MenuItem) error
	// DeleteMenuItem(ctx context.Context, id int) error
	GetAvailableTables(ctx context.Context) ([]models.Table, error)
}

type MenuUsecase struct {
	repo interfaces.Repository
}

func NewMenuUsecase(repo interfaces.Repository) MenuIUsecase {
	return &MenuUsecase{
		repo: repo,
	}
}

func (m *MenuUsecase) GetAllMenuItems(ctx context.Context) ([]models.MenuItem, error) {
	menuItems, err := m.repo.GetAllMenuItems(ctx)
	if err != nil {
		return nil, err
	}
	if len(menuItems) == 0 {
		return nil, models.ErrIsEmpty
	}
	return menuItems, nil
}

func (m *MenuUsecase) GetAvailableTables(ctx context.Context) ([]models.Table, error) {
	tables, err := m.repo.GetTablesByStatus(ctx, models.TableStatusAvailable)
	if err != nil {
		return nil, err
	}
	if len(tables) == 0 {
		return nil, models.ErrIsEmpty
	}
	return tables, nil
}
