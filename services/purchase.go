package services

import (
	"errors"
	"time"

	"avito.ru/shop/models"
	"avito.ru/shop/repositories"
)

// логика для покупок.
type PurchaseService struct {
	UserRepo     *repositories.UserRepository
	PurchaseRepo *repositories.PurchaseRepository
}

// Покупка товара
func (s *PurchaseService) BuyItem(user *models.User, itemName string) error {
	item, exists := models.MerchandiseList[itemName]
	if !exists {
		return errors.New("item not found in the shop")
	}

	if user.Balance < item.Price {
		return errors.New("insufficient balance")
	}

	// Вычитаем стоимость товара из баланса пользователя
	err := s.UserRepo.UpdateBalance(user.ID, -item.Price)
	if err != nil {
		return err
	}

	// Создаем запись о покупке
	purchase := &models.Purchase{
		UserID:      user.ID,
		ItemName:    item.Name,
		Price:       item.Price,
		PurchasedAt: time.Now(),
	}
	return s.PurchaseRepo.CreatePurchase(purchase)
}
