package services

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
	"time"
	"tiny-url/lib"
	"tiny-url/models"
	"tiny-url/repository"
)

// URLService service layer
type URLService struct {
	logger     lib.Logger
	repository repository.URLRepository
}

// NewURLService creates a new url service
func NewURLService(logger lib.Logger, repository repository.URLRepository) URLService {
	return URLService{
		logger:     logger,
		repository: repository,
	}
}

// WithTrx delegates transaction to repository database
func (s URLService) WithTrx(trxHandle *gorm.DB) URLService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// GetOneUrl gets one url
func (s URLService) GetOneUrl(hash string) (url *models.URL, err error) {
	err = s.repository.Find(&url, "hash = ?", hash).Error

	if err != nil {
		return
	}

	if time.Now().After(*url.ExpirationDate) {
		s.repository.Delete(url)
		return nil, nil
	}

	return
}

// GetAllUrl gets all url
func (s URLService) GetAllUrl() (urls []models.URL, err error) {
	return urls, s.repository.Find(&urls).Error
}

func (s URLService) CreateUrl(url *models.URL) (err error) {
	urlID, err := gonanoid.New(16)

	if err != nil {
		s.logger.Error(err)
		return err
	}

	url.Hash = &urlID
	now := time.Now()

	expiration := now.AddDate(2, 0, 0)

	url.CreationDate = &now
	url.ExpirationDate = &expiration

	return s.repository.Create(url).Error
}
