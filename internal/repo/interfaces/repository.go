package interfaces

import (
	"context"

	"github.com/abdullahnettoor/tastybites/internal/models"
)

type Repository interface {
	UserRepository
	OrderRepository
	MenuItemRepository
	TableRepository
}

// UserRepository defines user-related database operations.
type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GetUserById(ctx context.Context, id int) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, id int) error
}

// OrderRepository defines order-related database operations.
type OrderRepository interface {
	CreateOrder(ctx context.Context, order models.Order) (int, error)
	GetOrderById(ctx context.Context, id int) (models.Order, error)
	GetAllOrders(ctx context.Context) ([]models.Order, error)
	GetOrdersByUser(ctx context.Context, userId int) ([]models.Order, error)
	UpdateOrder(ctx context.Context, order models.Order) error
	DeleteOrder(ctx context.Context, id int) error
	GetOrderByTableId(ctx context.Context, tableId int) (models.Order, error)
	UpdateOrderStatusByTableId(ctx context.Context, tableId int, status models.OrderStatus) error
}

// MenuItemRepository defines menu item-related database operations.
type MenuItemRepository interface {
	CreateMenuItem(ctx context.Context, item models.MenuItem) (int, error)
	GetMenuItemById(ctx context.Context, id int) (models.MenuItem, error)
	GetMenuItemsByCategory(ctx context.Context, category string) ([]models.MenuItem, error)
	UpdateMenuItem(ctx context.Context, item models.MenuItem) error
	DeleteMenuItem(ctx context.Context, id int) error
	GetAllMenuItems(ctx context.Context) ([]models.MenuItem, error)
}

// TableRepository defines table-related database operations.
type TableRepository interface {
	CreateTable(ctx context.Context, table models.Table) (int, error)
	GetTableById(ctx context.Context, id int) (models.Table, error)
	UpdateTable(ctx context.Context, table models.Table) error
	DeleteTable(ctx context.Context, id int) error
	GetAllTables(ctx context.Context) ([]models.Table, error)
	GetTablesByStatus(ctx context.Context, status models.TableStatus) ([]models.Table, error)
	ResetTableToAvailable(ctx context.Context, tableId int) error
}
