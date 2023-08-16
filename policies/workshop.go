package policies

import (
	"athena_service/entities"
	"athena_service/utils"
	"context"
	"gorm.io/gorm"
)

type WorkshopPolicy struct {
	db *gorm.DB
}

func NewWorkshopPolicy(db *gorm.DB) WorkshopPolicy {
	return WorkshopPolicy{db: db}
}

func (p WorkshopPolicy) IsMember(ctx context.Context, workshopId int) (bool, error) {
	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		return false, err
	}

	var count int
	if err := p.db.Model(&entities.Member{}).
		Select("count(*)").
		Where("user_id = ? AND workshop_id = ?", user.Id, workshopId).
		Scan(&count).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (p WorkshopPolicy) IsTeacherInWorkshop(ctx context.Context, workshopId int) (bool, error) {
	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		return false, err
	}
	var member entities.Member
	err = p.db.Where("user_id = ? AND workshop_id = ?", user.Id, workshopId).First(&member).Error
	if err != nil {
		return false, err
	}
	return member.Role == "teacher", nil
}
