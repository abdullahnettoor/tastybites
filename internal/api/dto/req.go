package dto

import "github.com/abdullahnettoor/tastybites/internal/models"

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToUserModel(req any) *models.User {
	switch v := req.(type) {
	case UserRegisterRequest:
		return &models.User{
			Name:     v.Name,
			Email:    v.Email,
			Password: v.Password,
		}
	case UserLoginRequest:
		return &models.User{
			Email:    v.Email,
			Password: v.Password,
		}
	}
	return nil
}

type CreateOrderRequest struct {
	UserID  int         `json:"userId"`
	TableID int         `json:"tableId"`
	ItemsID []int       `json:"itemsId"`
	Items   []OrderItem `json:"items"`
}

type OrderItem struct {
	ItemID   int     `json:"itemId"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"` // Assuming price is provided in the request
}
