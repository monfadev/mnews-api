package news

import "gorm.io/gorm"

// / Repository layer
type Repository interface {
	FindAll() ([]News, error)
	FindByID(ID int) (News, error)
	Create(news News) (News, error)
	Update(news News) (News, error)
	Delete(news News) (News, error)
}

// / Object repository
type repository struct {
	db *gorm.DB
}

// / Initiate repository
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]News, error) {
	var news []News
	err := r.db.Find(&news).Error

	return news, err
}

func (r *repository) FindByID(ID int) (News, error) {
	var new News
	err := r.db.Find(&new, ID).Error

	return new, err
}

func (r *repository) Create(new News) (News, error) {
	err := r.db.Create(&new).Error

	return new, err
}

func (r *repository) Update(new News) (News, error) {
	err := r.db.Save(&new).Error

	return new, err
}

func (r *repository) Delete(new News) (News, error) {
	err := r.db.Delete(&new).Error

	return new, err
}
