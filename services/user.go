package services

import (
	"errors"

	"avito.ru/shop/models"
	"avito.ru/shop/repositories"
)

// логика для пользователей.

type UserService struct {
	UserRepo        *repositories.UserRepository
	TransactionRepo *repositories.TransactionRepository
}

// NewUserService создает новый экземпляр UserService
func NewUserService(userRepo *repositories.UserRepository, transactionRepo *repositories.TransactionRepository) *UserService {
	return &UserService{
		UserRepo:        userRepo,
		TransactionRepo: transactionRepo,
	}
}

// RegisterUser регистрирует нового пользователя
func (s *UserService) RegisterUser(username string) (*models.User, error) {
	existingUser, _ := s.UserRepo.GetUserByUsername(username)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	newUser := &models.User{
		Username: username,
		Balance:  1000,
	}
	err := s.UserRepo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}
	newUser, err = s.UserRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

// TransferCoins переводит монеты между пользователями
func (s *UserService) TransferCoins(fromUser, toUser *models.User, amount int) error {
	if fromUser.Balance < amount {
		return errors.New("insufficient balance")
	}

	err := s.UserRepo.UpdateBalance(fromUser.ID, -amount)
	if err != nil {
		return err
	}

	err = s.UserRepo.UpdateBalance(toUser.ID, amount)
	if err != nil {
		return err
	}

	txn := &models.Transaction{
		FromUserID:  fromUser.ID,
		ToUserID:    toUser.ID,
		Amount:      amount,
		Description: "Transfer",
	}
	return s.TransactionRepo.CreateTransaction(txn)
}

// GetUserByID получает пользователя по ID
func (s *UserService) GetUserByID(userID int64) (*models.User, error) {
	return s.UserRepo.GetUserByID(userID)
}
