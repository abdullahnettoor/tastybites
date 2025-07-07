package models

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusCompleted OrderStatus = "completed"
	OrderStatusCancelled OrderStatus = "cancelled"
)

func (os OrderStatus) String() string {
	return string(os)
}

type Order struct {
	ID         int         `json:"id"`
	UserID     int         `json:"userId"`
	Status     OrderStatus `json:"status"`  // e.g., "pending", "completed", "cancelled"
	ItemsID    []int       `json:"itemsId"` // List of item IDs in the order
	Items      []OrderItem `json:"items"`   // List of items in the order
	TotalPrice float64     `json:"totalPrice"`
	TableID    int         `json:"tableId"`
	CreatedAt  string      `json:"createdAt"`
	UpdatedAt  string      `json:"updatedAt"`
}

type OrderItem struct {
	MenuItemID int     `json:"menuItemId"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"` // Price per item
	// we can extend this with item level discounts or fields like cooking instructions etc...
}

func (o *Order) CalculateTotalPrice() {
	total := 0.0
	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}
	o.TotalPrice = total
}
