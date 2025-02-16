package repositories

import (
	"database/sql"

	"avito.ru/shop/models"
)

type InfoRepository struct {
	DB *sql.DB
}

func (r *InfoRepository) GetUserPurchases(userID int64) ([]models.Purchase, error) {
	rows, err := r.DB.Query("SELECT id, user_id, item_name, price FROM purchases WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var purchases []models.Purchase
	for rows.Next() {
		var p models.Purchase
		err := rows.Scan(&p.ID, &p.UserID, &p.ItemName, &p.Price)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, p)
	}
	return purchases, nil
}
