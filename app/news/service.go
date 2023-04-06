package news

type Service interface {
	FindAll() ([]News, error)
	FindByID(ID int) (News, error)
	Create(newsRequest NewsRequest) (News, error)
	Update(ID int, newsRequest NewsUpdateRequest) (News, error)
	Delete(ID int) (News, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]News, error) {
	news, err := s.repository.FindAll()
	return news, err
}

func (s *service) FindByID(ID int) (News, error) {
	return s.repository.FindByID(ID)
}

func (s *service) Create(newsRequest NewsRequest) (News, error) {
	new := News{
		Title:       newsRequest.Title,
		Description: newsRequest.Description,
		Author:      newsRequest.Author,
		PhoneNumber: newsRequest.PhoneNumber,
	}
	news, err := s.repository.Create(new)
	return news, err

}

func (s *service) Update(ID int, newsRequest NewsUpdateRequest) (News, error) {
	new, _ := s.repository.FindByID(ID)

	new.Title = newsRequest.Title
	new.Description = newsRequest.Description
	new.Author = newsRequest.Author
	new.PhoneNumber = newsRequest.PhoneNumber

	news, err := s.repository.Update(new)
	return news, err

}

func (s *service) Delete(ID int) (News, error) {
	new, _ := s.repository.FindByID(ID)
	news, err := s.repository.Delete(new)
	return news, err

}
