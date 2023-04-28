package repository

import "mnewsapi/app/models"

type Service interface {
	FindAll() ([]models.News, error)
	FindByID(ID int) (models.News, error)
	Create(newsRequest models.NewsRequest) (models.News, error)
	Update(ID int, newsRequest models.NewsUpdateRequest) (models.News, error)
	Delete(ID int) (models.News, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]models.News, error) {
	news, err := s.repository.FindAll()
	return news, err
}

func (s *service) FindByID(ID int) (models.News, error) {
	return s.repository.FindByID(ID)
}

func (s *service) Create(newsRequest models.NewsRequest) (models.News, error) {
	new := models.News{
		Title:       newsRequest.Title,
		Description: newsRequest.Description,
		Author:      newsRequest.Author,
		PhoneNumber: newsRequest.PhoneNumber,
	}
	news, err := s.repository.Create(new)
	return news, err

}

func (s *service) Update(ID int, newsRequest models.NewsUpdateRequest) (models.News, error) {
	new, _ := s.repository.FindByID(ID)

	new.Title = newsRequest.Title
	new.Description = newsRequest.Description
	new.Author = newsRequest.Author
	new.PhoneNumber = newsRequest.PhoneNumber

	news, err := s.repository.Update(new)
	return news, err

}

func (s *service) Delete(ID int) (models.News, error) {
	new, _ := s.repository.FindByID(ID)
	news, err := s.repository.Delete(new)
	return news, err

}
