package notification

import (
	"athena_service/base"
	"gorm.io/gorm"
)

type repo struct {
	base.Repository
}

func NewRepository(db *gorm.DB) repo {
	return repo{
		base.Repository{Db: db},
	}
}
