package pgrepo

import (
	"context"
	"fmt"

	"github.com/abdullahnettoor/tastybites/internal/models"
)

// Order operations
func (r *repository) CreateOrder(ctx context.Context, order models.Order) (int, error) {

	// Insert order into the database
	query := `INSERT INTO public.orders (user_id, table_id, total_price, status) VALUES ($1, $2, $3, $4) RETURNING id`
	var orderID int
	err := r.DB.QueryRowContext(ctx, query, order.UserID, order.TableID, order.TotalPrice, order.Status).Scan(&orderID)
	if err != nil {
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	// Insert order items into the database
	for _, item := range order.Items {
		itemQuery := `INSERT INTO public.order_items (order_id, menu_item_id, quantity, price) VALUES ($1, $2, $3, $4)`
		_, err := r.DB.ExecContext(ctx, itemQuery, orderID, item.MenuItemID, item.Quantity, item.Price)
		if err != nil {
			return 0, fmt.Errorf("failed to insert order item: %w", err)
		}
	}

	return orderID, nil
}

func (r *repository) GetOrderById(ctx context.Context, id int) (models.Order, error) {
	var order models.Order
	query := `SELECT id, user_id, total_price, status FROM public.orders WHERE id = $1`
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&order.ID, &order.UserID, &order.TotalPrice, &order.Status)
	if err != nil {
		return models.Order{}, fmt.Errorf("failed to get order by ID: %w", err)
	}

	// Get order items
	itemQuery := `SELECT menu_item_id, quantity, price FROM public.order_items WHERE order_id = $1`
	rows, err := r.DB.QueryContext(ctx, itemQuery, id)
	if err != nil {
		return models.Order{}, fmt.Errorf("failed to get order items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.OrderItem
		if err := rows.Scan(&item.MenuItemID, &item.Quantity, &item.Price); err != nil {
			return models.Order{}, fmt.Errorf("failed to scan order item: %w", err)
		}
		order.Items = append(order.Items, item)
	}

	return order, nil
}

func (r *repository) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	query := `SELECT id, user_id, total_price, status FROM public.orders`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all orders: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.TotalPrice, &order.Status); err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}

		// Get order items
		itemQuery := `SELECT menu_item_id, quantity, price FROM public.order_items WHERE order_id = $1`
		itemRows, err := r.DB.QueryContext(ctx, itemQuery, order.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get order items: %w", err)
		}
		defer itemRows.Close()

		for itemRows.Next() {
			var item models.OrderItem
			if err := itemRows.Scan(&item.MenuItemID, &item.Quantity, &item.Price); err != nil {
				return nil, fmt.Errorf("failed to scan order item: %w", err)
			}
			order.Items = append(order.Items, item)
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *repository) GetOrdersByUser(ctx context.Context, userId int) ([]models.Order, error) {
	var orders []models.Order
	query := `SELECT id, user_id, total_price, status FROM public.orders WHERE user_id = $1`
	rows, err := r.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders by user: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.TotalPrice, &order.Status); err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}
		order.Items = []models.OrderItem{}
		orders = append(orders, order)
	}

	// Fetch all orders and their items in a single query
	itemQuery := `
		SELECT oi.order_id, oi.menu_item_id, oi.quantity, oi.price
		FROM public.order_items oi
		JOIN public.orders o ON o.id = oi.order_id
		WHERE o.user_id = $1
	`
	itemRows, err := r.DB.QueryContext(ctx, itemQuery, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get order items: %w", err)
	}
	defer itemRows.Close()

	// Map orderID to order pointer for quick lookup
	orderMap := make(map[int]*models.Order)
	for i := range orders {
		orderMap[orders[i].ID] = &orders[i]
	}

	for itemRows.Next() {
		var orderID int
		var item models.OrderItem
		if err := itemRows.Scan(&orderID, &item.MenuItemID, &item.Quantity, &item.Price); err != nil {
			return nil, fmt.Errorf("failed to scan order item: %w", err)
		}
		if order, ok := orderMap[orderID]; ok {
			order.Items = append(order.Items, item)
		}
	}

	return orders, nil
}

func (r *repository) GetOrderByTableId(ctx context.Context, tableId int) (models.Order, error) {
	query := `
		SELECT o.id, o.user_id, o.total_price, o.status, oi.menu_item_id, oi.quantity, oi.price
		FROM public.orders o
		LEFT JOIN public.order_items oi ON o.id = oi.order_id
		WHERE o.table_id = $1 AND o.status = 'pending'
	`
	rows, err := r.DB.QueryContext(ctx, query, tableId)
	if err != nil {
		return models.Order{}, fmt.Errorf("failed to get order by table ID: %w", err)
	}
	defer rows.Close()

	var order models.Order
	order.Items = []models.OrderItem{}
	first := true
	for rows.Next() {
		var (
			id         int
			userID     int
			totalPrice float64
			status     string
			menuItemID *int
			quantity   *int
			price      *float64
		)
		if err := rows.Scan(&id, &userID, &totalPrice, &status, &menuItemID, &quantity, &price); err != nil {
			return models.Order{}, fmt.Errorf("failed to scan row: %w", err)
		}
		if first {
			order.ID = id
			order.UserID = userID
			order.TableID = tableId
			order.Status = models.OrderStatus(status)
			first = false
		}

		if menuItemID != nil && quantity != nil && price != nil {
			order.Items = append(order.Items, models.OrderItem{
				MenuItemID: *menuItemID,
				Quantity:   *quantity,
				Price:      *price,
			})
		}
	}

	if first {
		return models.Order{}, fmt.Errorf("no pending order found for table ID %d", tableId)
	}

	return order, nil
}

func (r *repository) UpdateOrder(ctx context.Context, order models.Order) error {
	return fmt.Errorf("not implemented")
}

func (r *repository) DeleteOrder(ctx context.Context, id int) error {
	return fmt.Errorf("not implemented")
}

func (r *repository) UpdateOrderStatusByTableId(ctx context.Context, tableId int, status models.OrderStatus) error {
	query := `UPDATE public.orders SET status = $1 WHERE table_id = $2 AND status = 'pending'`
	_, err := r.DB.ExecContext(ctx, query, status, tableId)
	if err != nil {
		return fmt.Errorf("failed to update order status by table ID: %w", err)
	}
	return nil
}
