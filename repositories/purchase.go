package repositories

import (
	"database/sql"
	"time"

	"avito.ru/shop/models"
)

// Репозиторий для работы с таблицей purchases.
type PurchaseRepository struct {
	DB *sql.DB
}

// Создание записи о покупке
func (r *PurchaseRepository) CreatePurchase(purchase *models.Purchase) error {
	_, err := r.DB.Exec("INSERT INTO purchases (user_id, item_name, price, purchased_at) VALUES ($1, $2, $3, $4)",
		purchase.UserID, purchase.ItemName, purchase.Price, time.Now())
	return err
}

// Получение списка покупок пользователя
func (r *PurchaseRepository) GetUserPurchases(userID int64) ([]models.Purchase, error) {
	rows, err := r.DB.Query("SELECT id, user_id, item_name, price, purchased_at FROM purchases WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var purchases []models.Purchase
	for rows.Next() {
		var p models.Purchase
		err := rows.Scan(&p.ID, &p.UserID, &p.ItemName, &p.Price, &p.PurchasedAt)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, p)
	}
	return purchases, nil
}
