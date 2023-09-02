package notification

import (
	"athena_service/app"
	"athena_service/dto"
	"athena_service/entities"
	"athena_service/utils"
	"context"
)

type IPolicy interface {
	IsTeacherInWorkshop(ctx context.Context, workshopID int) (bool, error)
	IsMember(ctx context.Context, workshopId int) (bool, error)
}

type usecase struct {
	repository repo
	policy     IPolicy
}

func NewUsecase(repo repo, policy IPolicy) usecase {
	return usecase{repository: repo, policy: policy}
}

func (u usecase) CreateNotificationWorkshop(ctx context.Context, input dto.CreateNotificationWorkshopInput) error {
	isTeacher, err := u.policy.IsTeacherInWorkshop(ctx, input.WorkshopId)
	if err != nil || !isTeacher {
		return app.NewForbiddenError("forbidden").WithError(err)
	}

	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		return app.NewForbiddenError("forbidden").WithError(err)
	}

	notification := entities.NotificationWorkshop{
		WorkshopId: input.WorkshopId,
		Content:    input.Content,
		UserId:     user.Id,
	}
	err = u.repository.GetDB(ctx).Create(&notification).Error
	if err != nil {
		return err
	}

	return nil

}

func (u usecase) GetNotificationWorkshop(ctx context.Context, workshopId int) ([]entities.NotificationWorkshop, error) {
	check, err := u.policy.IsMember(ctx, workshopId)
	if err != nil || !check {
		return nil, app.NewForbiddenError("forbidden").WithError(err)
	}
	var notifications []entities.NotificationWorkshop
	err = u.repository.GetDB(ctx).Where("workshop_id = ?", workshopId).Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}
