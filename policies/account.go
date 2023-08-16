package policies

import (
	"athena_service/utils"
	"context"
	"gorm.io/gorm"
)

type AccountPolicy struct {
	Db *gorm.DB
}

func NewAccountPolicy(db *gorm.DB) AccountPolicy {
	return AccountPolicy{Db: db}
}

func (p AccountPolicy) IsTeacher(ctx context.Context) (bool, error) {
	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		return false, err
	}
	return user.Role == "teacher", nil
}

func (p AccountPolicy) IsStudent(ctx context.Context) (bool, error) {
	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		return false, err
	}
	return user.Role == "student", nil
}
