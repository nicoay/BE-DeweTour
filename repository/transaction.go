package repository

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(trans models.Transaction) (models.Transaction, error)
	UpdateTransaction(trans models.Transaction) (models.Transaction, error)
	DeleteTransaction(ID int, trans models.Transaction) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var Transactions []models.Transaction
	err := r.db.Preload("User").Preload("Tour.Countries").Find(&Transactions).Error

	return Transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var Transaction models.Transaction
	err := r.db.Preload("User").Preload("Tour.Countries").First(&Transaction, ID).Error

	return Transaction, err
}

func (r *repository) GetCountryTransaction(ID int) (models.CountryResponse, error) {
	var Country models.CountryResponse
	err := r.db.First(&Country, ID).Error

	return Country, err
}

func (r *repository) CreateTransaction(Transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Preload("Tour.Countries").Create(&Transaction).Error

	return Transaction, err
}

func (r *repository) UpdateTransaction(trans models.Transaction) (models.Transaction, error) {
	err := r.db.Save(&trans).Error
	return trans, err
}

func (r *repository) DeleteTransaction(ID int, Transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&Transaction).Error
	return Transaction, err
}
