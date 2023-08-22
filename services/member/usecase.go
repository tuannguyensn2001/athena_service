package member

import (
	"athena_service/app"
	"athena_service/constant"
	"athena_service/dto"
	"athena_service/entities"
	"athena_service/utils"
	"context"
	"errors"
	"gorm.io/gorm"
)

type usecase struct {
	repository repo
	policy     IPolicy
}

type IPolicy interface {
	IsTeacherInWorkshop(ctx context.Context, workshopId int) (bool, error)
	IsStudentWithId(ctx context.Context, id int) (bool, error)
	IsStudent(ctx context.Context) (bool, error)
}

func NewUsecase(repo repo, policy IPolicy) usecase {
	return usecase{repository: repo, policy: policy}
}

func (u usecase) GetStudent(ctx context.Context, workshopId int) ([]entities.User, error) {
	isTeacher, err := u.policy.IsTeacherInWorkshop(ctx, workshopId)
	if err != nil || !isTeacher {
		return nil, app.NewForbiddenError("forbidden").WithError(err)
	}

	var ids []int

	if err := u.repository.GetDB(ctx).Model(&entities.Member{}).Select("user_id").Where("workshop_id = ? and role = ? and status = ?", workshopId, constant.STUDENT, constant.ACTIVE).Find(&ids).Error; err != nil {
		return nil, err
	}

	var users []entities.User
	if err := u.repository.GetDB(ctx).Preload("Profile").Where("id in ?", ids).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil

}

func (u usecase) AddStudent(ctx context.Context, input dto.AddStudentInput) error {

	var user entities.User
	if err := u.repository.GetDB(ctx).Where("phone = ?", input.Phone).First(&user).Error; err != nil {
		return err
	}

	isTeacher, err := u.policy.IsTeacherInWorkshop(ctx, input.WorkshopId)
	if err != nil || !isTeacher {
		return app.NewForbiddenError("forbidden").WithError(err)
	}
	isStudent, err := u.policy.IsStudentWithId(ctx, user.Id)
	if err != nil || !isStudent {
		return app.NewForbiddenError("forbidden").WithError(err)
	}

	var count int64
	if err := u.repository.GetDB(ctx).Model(&entities.Member{}).
		Where("workshop_id = ? and user_id = ? and role = ?", input.WorkshopId, user.Id, constant.STUDENT).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return app.NewBadRequestError("student already exist")
	}

	member := entities.Member{
		WorkshopId: input.WorkshopId,
		UserId:     user.Id,
		Role:       constant.STUDENT,
		Status:     constant.ACTIVE,
	}

	if err := u.repository.GetDB(ctx).Create(&member).Error; err != nil {
		return err
	}

	return nil

}

func (u usecase) StudentRequestJoin(ctx context.Context, input dto.StudentRequestJoinInput) error {
	isStudent, err := u.policy.IsStudent(ctx)
	if err != nil || !isStudent {
		return app.NewForbiddenError("forbidden").WithError(err)
	}
	var workshop entities.Workshop
	if err := u.repository.GetDB(ctx).Where("id = ?", input.WorkshopId).First(&workshop).Error; err != nil {
		return err
	}

	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		return err
	}

	if workshop.IsLock {
		return app.NewBadRequestError("workshop is lock")
	}

	var member entities.Member
	err = u.repository.GetDB(ctx).Where("workshop_id = ? and user_id = ?", input.WorkshopId, user.Id).First(&member).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err == nil {
		if member.Status == constant.ACTIVE {
			return app.NewBadRequestError("student already join")
		}
		if member.Status == constant.PENDING {
			return app.NewBadRequestError("student already request")
		}
	}

	member = entities.Member{
		WorkshopId: input.WorkshopId,
		UserId:     user.Id,
		Role:       constant.STUDENT,
	}

	if !workshop.ApproveStudent {
		member.Status = constant.ACTIVE
	} else {
		member.Status = constant.PENDING
	}

	if err := u.repository.GetDB(ctx).Create(&member).Error; err != nil {
		return err
	}

	return nil

}
