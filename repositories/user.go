package repositories

import (
	"database/sql"
	"errors"

	"avito.ru/shop/models"
)

type UserRepository struct {
	DB *sql.DB
}

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

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.DB.Exec("INSERT INTO users (username, balance) VALUES ($1, $2)", user.Username, user.Balance)
	return err
}

func (r *UserRepository) UpdateBalance(userID int64, amount int) error {
	_, err := r.DB.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", amount, userID)
	return err
}
