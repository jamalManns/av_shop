package repositories

import (
	"database/sql"

	"avito.ru/shop/models"
)

type TransactionRepository struct {
	DB *sql.DB
}

func (r *TransactionRepository) CreateTransaction(txn *models.Transaction) error {
	_, err := r.DB.Exec("INSERT INTO transactions (from_user_id, to_user_id, amount, description) VALUES ($1, $2, $3, $4)",
		txn.FromUserID, txn.ToUserID, txn.Amount, txn.Description)
	return err
}

func (r *TransactionRepository) GetTransactionsByUserID(userID int64) ([]models.Transaction, error) {
	rows, err := r.DB.Query("SELECT id, from_user_id, to_user_id, amount, description FROM transactions WHERE from_user_id = $1 OR to_user_id = $2", userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var txn models.Transaction
		err := rows.Scan(&txn.ID, &txn.FromUserID, &txn.ToUserID, &txn.Amount, &txn.Description)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, txn)
	}
	return transactions, nil
}
