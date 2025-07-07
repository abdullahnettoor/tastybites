package models

type MenuItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	ImageURL    string `json:"imageUrl"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`

	// this can be extended with available options, stock management, etc.
}