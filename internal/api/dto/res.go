package dto

type CommonResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  interface{} `json:"user"`
}