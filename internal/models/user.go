package models

type UserRole string

const (
	UserRoleAdmin   UserRole = "admin"
	UserRoleUser    UserRole = "user"
	UserRoleManager UserRole = "manager"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Role      string `json:"role"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
