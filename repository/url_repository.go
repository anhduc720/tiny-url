package repository

import (
	"gorm.io/gorm"
	"tiny-url/lib"
)

type URLRepository struct {
	lib.Database
	logger lib.Logger
}

func NewURLRepository(db lib.Database, logger lib.Logger) URLRepository {
	return URLRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx enables repository with transaction
func (r URLRepository) WithTrx(trxHandle *gorm.DB) URLRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
