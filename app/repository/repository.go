package repository

import (
	"mnewsapi/app/models"

	"gorm.io/gorm"
)

// / Repository layer
type Repository interface {
	FindAll() ([]models.News, error)
	FindByID(ID int) (models.News, error)
	Create(news models.News) (models.News, error)
	Update(news models.News) (models.News, error)
	Delete(news models.News) (models.News, error)
}

// / Object repository
type repository struct {
	db *gorm.DB
}

// / Initiate repository
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.News, error) {
	var news []models.News
	err := r.db.Find(&news).Error

	return news, err
}

func (r *repository) FindByID(ID int) (models.News, error) {
	var new models.News
	err := r.db.Find(&new, ID).Error

	return new, err
}

func (r *repository) Create(new models.News) (models.News, error) {
	err := r.db.Create(&new).Error

	return new, err
}

func (r *repository) Update(new models.News) (models.News, error) {
	err := r.db.Save(&new).Error

	return new, err
}

func (r *repository) Delete(new models.News) (models.News, error) {
	err := r.db.Delete(&new).Error

	return new, err
}
