package workshop

import (
	"athena_service/app"
	"athena_service/constant"
	"athena_service/dto"
	"athena_service/entities"
	"athena_service/utils"
	"context"
)

type usecase struct {
	repository repo
	policy     IPolicy
}

type IPolicy interface {
	IsTeacher(ctx context.Context) (bool, error)
}

func NewUsecase(repository repo, policy IPolicy) usecase {
	return usecase{repository: repository, policy: policy}
}

func (u usecase) Create(ctx context.Context, input dto.CreateWorkshopInput) error {
	isTeacher, err := u.policy.IsTeacher(ctx)
	if err != nil {
		return err
	}
	if !isTeacher {
		return app.NewForbiddenError("forbidden")
	}

	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		return err
	}

	code := utils.RandomUppercase(5)

	err = u.repository.Transaction(ctx, func(ctx context.Context) error {

		thumbnail := input.Thumbnail
		if len(thumbnail) == 0 {
			thumbnail = "https://shub-storage.sgp1.cdn.digitaloceanspaces.com/profile_images/44-01.jpg"
		}

		workshop := entities.Workshop{
			Name:                input.Name,
			Thumbnail:           thumbnail,
			PrivateCode:         input.PrivateCode,
			Code:                code,
			ApproveStudent:      input.ApproveStudent,
			PreventStudentLeave: input.PreventStudentLeave,
			ApproveShowScore:    input.ApproveShowScore,
			DisableNewsfeed:     input.DisableNewsfeed,
			LimitPolicyTeacher:  input.LimitPolicyTeacher,
			IsShow:              true,
			Subject:             input.Subject,
			Grade:               input.Grade,
			IsLock:              false,
		}

		err := u.repository.Create(ctx, &workshop)
		if err != nil {
			return err
		}
		member := entities.Member{
			WorkshopId: workshop.Id,
			UserId:     user.Id,
			Role:       "teacher",
			Status:     constant.ACTIVE,
		}
		err = u.repository.Create(ctx, &member)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (u usecase) GetOwn(ctx context.Context, input dto.GetOwnWorkshopInput) ([]entities.Workshop, error) {
	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var workshopIds []int
	if err := u.repository.GetDB(ctx).Model(&entities.Member{}).Select("workshop_id").Where("user_id = ? and role = ?", user.Id, user.Role).Find(&workshopIds).Error; err != nil {
		return nil, err
	}

	var workshops []entities.Workshop
	if err := u.repository.GetDB(ctx).Where("id in ?", workshopIds).Where("is_show = ?", input.IsShow).Find(&workshops).Error; err != nil {
		return nil, err
	}

	return workshops, nil

}

func (u usecase) FindByCode(ctx context.Context, code string) (entities.Workshop, error) {
	var result entities.Workshop

	if err := u.repository.GetDB(ctx).Where("code = ?", code).First(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
