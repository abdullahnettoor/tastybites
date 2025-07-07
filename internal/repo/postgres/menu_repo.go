package pgrepo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/abdullahnettoor/tastybites/internal/models"
)

// Menu operations
func (r *repository) CreateMenuItem(ctx context.Context, item models.MenuItem) (int, error) {
	return 0, fmt.Errorf("not implemented")
}

func (r *repository) GetMenuItemById(ctx context.Context, id int) (models.MenuItem, error) {
	query := `SELECT id, name, description, price, category, image_url FROM public.menu_items WHERE id = $1`
	var item models.MenuItem
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Category, &item.ImageURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.MenuItem{}, fmt.Errorf("menu item not found with id %d: %w", id, err)
		}
		return models.MenuItem{}, fmt.Errorf("failed to get menu item by ID: %w", err)
	}
	return item, nil
}

func (r *repository) GetMenuItemsByCategory(ctx context.Context, category string) ([]models.MenuItem, error) {
	query := `SELECT id, name, description, price, category, image_url FROM public.menu_items WHERE category = $1`
	rows, err := r.DB.QueryContext(ctx, query, category)
	if err != nil {
		return nil, fmt.Errorf("failed to get menu items by category %s: %w", category, err)
	}
	defer rows.Close()

	var menuItems = make([]models.MenuItem, 0)
	for rows.Next() {
		var item models.MenuItem
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Category, &item.ImageURL); err != nil {
			return nil, fmt.Errorf("failed to scan menu item: %w", err)
		}
		menuItems = append(menuItems, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over menu items: %w", err)
	}
	if len(menuItems) == 0 {
		return nil, fmt.Errorf("no menu items found")
	}
	return menuItems, nil
}

func (r *repository) UpdateMenuItem(ctx context.Context, item models.MenuItem) error {
	return fmt.Errorf("not implemented")
}

func (r *repository) DeleteMenuItem(ctx context.Context, id int) error {
	return fmt.Errorf("not implemented")
}

func (r *repository) GetAllMenuItems(ctx context.Context) ([]models.MenuItem, error) {

	query := `SELECT id, name, description, price, category, image_url FROM public.menu_items`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all menu items: %w", err)
	}
	defer rows.Close()

	var menuItems = make([]models.MenuItem, 0)
	for rows.Next() {
		var item models.MenuItem
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Category, &item.ImageURL); err != nil {
			return nil, fmt.Errorf("failed to scan menu item: %w", err)
		}
		menuItems = append(menuItems, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over menu items: %w", err)
	}
	if len(menuItems) == 0 {
		return nil, fmt.Errorf("no menu items found")
	}
	return menuItems, nil
}
