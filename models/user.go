package models

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Balance  int    `json:"balance"`
}

// AuthRequest содержит данные для аутентификации
type AuthRequest struct {
	Username string `json:"username" binding:"required"`
}

// ErrorResponse используется для ответов об ошибках
type ErrorResponse struct {
	Errors string `json:"errors"`
}

// AuthResponse содержит токен после успешной аутентификации
type AuthResponse struct {
	Token string `json:"token"`
}

// SendCoinRequest содержит данные для перевода монет
type SendCoinRequest struct {
	ToUser string `json:"toUser" binding:"required"`
	Amount int    `json:"amount" binding:"required,gt=0"`
}
