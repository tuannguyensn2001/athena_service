package auth

import (
	"athena_service/base"
	"athena_service/entities"
	"context"
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

func (r *repo) FindByPhone(ctx context.Context, phone string) (*entities.User, error) {
	var result entities.User
	if err := r.GetDB(ctx).WithContext(ctx).Where("phone = ?", phone).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
