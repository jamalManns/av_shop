package services

import (
	"avito.ru/shop/models"
	"avito.ru/shop/repositories"
)

type InfoService struct {
	UserRepo        *repositories.UserRepository
	TransactionRepo *repositories.TransactionRepository
	PurchaseRepo    *repositories.PurchaseRepository
}

func NewInfoService(userRepo *repositories.UserRepository, transactionRepo *repositories.TransactionRepository, purchaseRepo *repositories.PurchaseRepository) *InfoService {
	return &InfoService{
		UserRepo:        userRepo,
		TransactionRepo: transactionRepo,
		PurchaseRepo:    purchaseRepo,
	}
}

func (s *InfoService) GetUserInfo(userID int64) (*models.InfoResponse, error) {
	// Получаем пользователя
	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Получаем историю транзакций
	transactions, err := s.TransactionRepo.GetTransactionsByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Получаем список покупок
	purchases, err := s.PurchaseRepo.GetUserPurchases(userID)
	if err != nil {
		return nil, err
	}

	// Формируем ответ
	response := &models.InfoResponse{
		Coins:     user.Balance,
		Inventory: groupPurchases(purchases),
		CoinHistory: models.CoinHist{
			Received: filterTransactions(transactions, userID, true),
			Sent:     filterTransactions(transactions, userID, false),
		},
	}

	return response, nil
}

// Группировка покупок по типам товаров
func groupPurchases(purchases []models.Purchase) []models.Item {
	inventory := make(map[string]int)
	for _, purchase := range purchases {
		inventory[purchase.ItemName]++
	}

	var items []models.Item
	for itemName, quantity := range inventory {
		items = append(items, models.Item{
			Type:     itemName,
			Quantity: quantity,
		})
	}

	return items
}

// Фильтрация транзакций
func filterTransactions(transactions []models.Transaction, userID int64, isReceived bool) []models.Transaction {
	var result []models.Transaction
	for _, txn := range transactions {
		if isReceived && txn.ToUserID == userID {
			result = append(result, txn)
		} else if !isReceived && txn.FromUserID == userID {
			result = append(result, txn)
		}
	}
	return result
}
