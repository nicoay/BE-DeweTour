package repository

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUserById(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(ID int, user models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var Users []models.User
	err := r.db.Preload("Transaction.Tour.Countries").Find(&Users).Error

	return Users, err
}

func (r *repository) GetUserById(ID int) (models.User, error) {
	var User models.User
	err := r.db.Preload("Transaction.Tour.Countries").First(&User, ID).Error
	return User, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Preload("Transaction.Tour.Countries").Create(&user).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Preload("Transaction.Tour.Countries").Save(&user).Error
	return user, err
}

func (r *repository) DeleteUser(ID int, user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error
	return user, err
}
