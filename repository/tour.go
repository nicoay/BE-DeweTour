package repository

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type TourRepository interface {
	FindTours() ([]models.Tour, error)
	GetTour(ID int) (models.Tour, error)
	CreateTour(tour models.Tour) (models.Tour, error)
	UpdateTour(tour models.Tour) (models.Tour, error)
	DeleteTour(ID int, tour models.Tour) (models.Tour, error)
	GetCountryTour(ID int) (models.CountryResponse, error)
}

func RepositoryTour(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTours() ([]models.Tour, error) {
	var tours []models.Tour
	err := r.db.Preload("Countries").Find(&tours).Error

	return tours, err
}

func (r *repository) GetTour(ID int) (models.Tour, error) {
	var Tour models.Tour
	err := r.db.Preload("Countries").First(&Tour, ID).Error

	return Tour, err
}

func (r *repository) GetCountryTour(ID int) (models.CountryResponse, error) {
	var Country models.CountryResponse
	err := r.db.First(&Country, ID).Error

	return Country, err
}

func (r *repository) CreateTour(tour models.Tour) (models.Tour, error) {
	err := r.db.Preload("Countries").Create(&tour).Error

	return tour, err
}
func (r *repository) UpdateTour(Tour models.Tour) (models.Tour, error) {
	err := r.db.Preload("Countries").Save(&Tour).Error
	return Tour, err
}

func (r *repository) DeleteTour(ID int, Tour models.Tour) (models.Tour, error) {
	err := r.db.Delete(&Tour).Error
	return Tour, err
}
