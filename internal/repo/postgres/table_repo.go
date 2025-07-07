package pgrepo

import (
	"context"
	"fmt"

	"github.com/abdullahnettoor/tastybites/internal/models"
)

// Table operations
func (r *repository) CreateTable(ctx context.Context, table models.Table) (int, error) {

	query := `INSERT INTO public.tables (name, seats, status) VALUES ($1, $2, $3) RETURNING id`
	var tableID int
	err := r.DB.QueryRowContext(ctx, query, table.Name, table.Seats, table.Status).Scan(&tableID)
	if err != nil {
		return 0, fmt.Errorf("failed to create table: %w", err)
	}
	return tableID, nil
}

func (r *repository) GetTableByName(ctx context.Context, name string) (models.Table, error) {
	var table models.Table
	query := `SELECT id, name, seats, status FROM public.tables WHERE name = $1`
	err := r.DB.QueryRowContext(ctx, query, name).Scan(&table.ID, &table.Name, &table.Seats, &table.Status)
	if err != nil {
		return models.Table{}, fmt.Errorf("failed to get table by name: %w", err)
	}
	return table, nil
}

func (r *repository) GetTableById(ctx context.Context, id int) (models.Table, error) {
	var table models.Table
	query := `SELECT id, name, seats, status FROM public.tables WHERE id = $1`
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&table.ID, &table.Name, &table.Seats, &table.Status)
	if err != nil {
		return models.Table{}, fmt.Errorf("failed to get table by ID: %w", err)
	}
	return table, nil
}

func (r *repository) UpdateTable(ctx context.Context, table models.Table) error {
	query := `UPDATE public.tables SET name = $1, seats = $2, status = $3 WHERE id = $4`
	_, err := r.DB.ExecContext(ctx, query, table.Name, table.Seats, table.Status, table.ID)
	if err != nil {
		return fmt.Errorf("failed to update table: %w", err)
	}
	return nil
}

func (r *repository) GetTablesByStatus(ctx context.Context, status models.TableStatus) ([]models.Table, error) {
	query := `SELECT id, name, seats, status FROM public.tables WHERE status = $1`
	rows, err := r.DB.QueryContext(ctx, query, status)
	if err != nil {
		return nil, fmt.Errorf("failed to get tables by status: %w", err)
	}
	defer rows.Close()

	var tables []models.Table
	for rows.Next() {
		var table models.Table
		if err := rows.Scan(&table.ID, &table.Name, &table.Seats, &table.Status); err != nil {
			return nil, fmt.Errorf("failed to scan table row: %w", err)
		}
		tables = append(tables, table)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over table rows: %w", err)
	}
	return tables, nil
}

func (r *repository) DeleteTable(ctx context.Context, id int) error {
	query := `DELETE FROM public.tables WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete table: %w", err)
	}
	return nil
}

func (r *repository) GetAllTables(ctx context.Context) ([]models.Table, error) {
	query := `SELECT id, name, capacity, status FROM public.tables`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all tables: %w", err)
	}
	defer rows.Close()

	var tables []models.Table
	for rows.Next() {
		var table models.Table
		if err := rows.Scan(&table.ID, &table.Name, &table.Seats, &table.Status); err != nil {
			return nil, fmt.Errorf("failed to scan table row: %w", err)
		}
		tables = append(tables, table)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over table rows: %w", err)
	}
	return tables, nil
}
