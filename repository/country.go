package repository

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type CountryRepository interface {
	FindCountries() ([]models.Country, error)
	GetCountry(ID int) (models.Country, error)
	CreateCountry(country models.Country) (models.Country, error)
	UpdateCountry(country models.Country) (models.Country, error)
	// DeleteUser(ID int, user models.User) (models.User, error)
}

func RepositoryCountry(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCountries() ([]models.Country, error) {
	var Countries []models.Country
	err := r.db.Find(&Countries).Error

	return Countries, err
}

func (r *repository) GetCountry(ID int) (models.Country, error) {
	var Country models.Country
	err := r.db.First(&Country, ID).Error
	return Country, err
}

func (r *repository) CreateCountry(country models.Country) (models.Country, error) {
	err := r.db.Create(&country).Error

	return country, err
}

func (r *repository) UpdateCountry(country models.Country) (models.Country, error) {
	err := r.db.Save(&country).Error
	return country, err
}

// func (r *repository) DeleteUser(ID int, user models.User) (models.User, error) {
// 	err := r.db.Delete(&user).Error
// 	return user, err
// }
