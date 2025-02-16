package repositories

import (
	"database/sql"
	"errors"

	"avito.ru/shop/models"
)

type UserRepository struct {
	DB *sql.DB
}

// GetUserByUsername получает пользователя по имени
func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT id, username, balance FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// CreateUser создает нового пользователя
func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.DB.Exec("INSERT INTO users (username, balance) VALUES ($1, $2)", user.Username, user.Balance)
	return err
}

// UpdateBalance обновляет баланс пользователя
func (r *UserRepository) UpdateBalance(userID int64, amount int) error {
	_, err := r.DB.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", amount, userID)
	return err
}

// GetUserByID получает пользователя по ID
func (r *UserRepository) GetUserByID(userID int64) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT id, username, balance FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Username, &user.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
