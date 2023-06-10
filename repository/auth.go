package repository

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(Email string) (models.User,error)
	CheckAuth(ID int) (models.User,error)
}

func RepositoryAuth(db * gorm.DB) *repository{
	return &repository{db}
}

func (r* repository) Register(user models.User)(models.User, error){
	err := r.db.Create(&user).Error

	return user,err
}

func (r *repository) Login(Email string)(models.User,error){
	var user models.User
	// "email=? untuk kondisi where di gorm"
	err := r.db.First(&user,"email=?" , Email).Error

	return user,err
}

func (r *repository) CheckAuth(ID int) (models.User, error) {
	var User models.User
	err := r.db.Preload("Transaction.Tour.Countries").First(&User, ID).Error
	return User, err
}