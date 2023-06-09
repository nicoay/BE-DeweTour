package repository

import (
	"dumbmerch/models"
	"fmt"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(trans models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, transId int) (models.Transaction, error)
	DeleteTransaction(ID int, trans models.Transaction) (models.Transaction, error)
	GetTourById(ID int) (models.TourResponse, error)
	GetUser(ID int) (models.UsersProfileResponse, error)
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

func (r *repository) GetUser(ID int) (models.UsersProfileResponse, error) {
	var User models.UsersProfileResponse
	err := r.db.First(&User, ID).Error

	return User, err
}
func (r *repository) GetTourById(ID int) (models.TourResponse, error) {
	var Tour models.TourResponse
	err := r.db.Preload("Countries").First(&Tour, ID).Error

	return Tour, err
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

func (r *repository) UpdateTransaction(status string, transId int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("User").Preload("Tour.Countries").First(&transaction, transId)
	fmt.Println(transaction.Tour.ID, "ini transaction")
	if status != transaction.Status && status == "success" {
		var tour models.Tour
		r.db.First(&tour, transaction.TourID)
		tour.QuotaCurrent = tour.QuotaCurrent + transaction.CounterQty
		r.db.Save(&tour)
	}

	transaction.Status = status
	err := r.db.Save(&transaction).Error

	return transaction, err

}

func (r *repository) DeleteTransaction(ID int, Transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&Transaction).Error
	return Transaction, err
}
