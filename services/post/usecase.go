package post

import (
	"athena_service/app"
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
	IsMember(ctx context.Context, workshopId int) (bool, error)
}

func NewUsecase(repo repo, policy IPolicy) usecase {
	return usecase{repository: repo, policy: policy}
}

func (u usecase) Create(ctx context.Context, input dto.CreatePostInput) error {
	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		return err
	}

	isMember, err := u.policy.IsMember(ctx, input.WorkshopId)
	if err != nil {
		return err
	}
	if !isMember {
		return app.NewForbiddenError("forbidden")
	}

	post := entities.Post{
		Content:    input.Content,
		UserId:     user.Id,
		WorkshopId: input.WorkshopId,
	}
	err = u.repository.GetDB(ctx).Create(&post).Error
	if err != nil {
		return err
	}

	return nil

}

func (u usecase) GetInWorkshop(ctx context.Context, workshopId int) ([]entities.Post, error) {
	isMember, err := u.policy.IsMember(ctx, workshopId)
	if err != nil || !isMember {
		return nil, app.NewForbiddenError("forbidden").WithError(err)
	}

	var posts []entities.Post

	if err := u.repository.GetDB(ctx).Preload("Comments").Preload("User").Preload("User.Profile").Where("workshop_id = ?", workshopId).
		Order("created_at desc").Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil

}
