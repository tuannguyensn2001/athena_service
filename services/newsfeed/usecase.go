package newsfeed

import (
	"athena_service/app"
	"athena_service/dto"
	"athena_service/entities"
	"athena_service/utils"
	"context"
	"fmt"
	"github.com/pusher/pusher-http-go/v5"
)

type usecase struct {
	repository repo
	policy     IPolicy
	pusher     *pusher.Client
}
type IPolicy interface {
	IsMember(ctx context.Context, workshopId int) (bool, error)
}

func NewUsecase(repo repo, policy IPolicy, pusher *pusher.Client) usecase {
	return usecase{repository: repo, policy: policy, pusher: pusher}
}

func (u usecase) CreatePost(ctx context.Context, input dto.CreatePostInput) error {
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

	var workshop entities.Workshop
	if err := u.repository.GetDB(ctx).Where("id = ?", input.WorkshopId).First(&workshop).Error; err != nil {
		return err
	}

	pusherClient := u.pusher
	err = pusherClient.Trigger(fmt.Sprintf("newsfeed-workshop-%s", workshop.Code), "new-post", post)
	if err != nil {
		return err
	}

	return nil

}

func (u usecase) GetPostsInWorkshop(ctx context.Context, input dto.GetPostInWorkshopInput) (dto.GetPostInWorkshopOutput, error) {
	var result dto.GetPostInWorkshopOutput
	isMember, err := u.policy.IsMember(ctx, input.WorkshopId)
	if err != nil || !isMember {
		return result, app.NewForbiddenError("forbidden").WithError(err)
	}

	var posts []entities.Post

	limit := 3
	page := input.Page

	if err := u.repository.GetDB(ctx).Preload("User").Preload("User.Profile").Where("workshop_id = ?", input.WorkshopId).
		Limit(limit).Offset((page - 1) * limit).
		Order("created_at desc").Find(&posts).Error; err != nil {
		return result, err
	}
	var count int64
	if err := u.repository.GetDB(ctx).Model(&entities.Post{}).Where("workshop_id = ?", input.WorkshopId).Count(&count).Error; err != nil {
		return result, err
	}

	result.Data = posts
	result.Meta.Total = int(count)
	result.Meta.Page = page

	return result, nil
}

func (u usecase) GetWorkshopByPostId(ctx context.Context, postId int) (entities.Workshop, error) {
	var workshop entities.Workshop
	if err := u.repository.GetDB(ctx).
		Raw("select * from workshops left join posts on workshops.id = posts.workshop_id where posts.id = ?", postId).
		Scan(&workshop).Error; err != nil {
		return workshop, err
	}
	return workshop, nil

}

func (u usecase) GetCommentsInPosts(ctx context.Context, postId int) ([]entities.Comment, error) {

	workshop, err := u.GetWorkshopByPostId(ctx, postId)
	if err != nil {
		return nil, err
	}
	isMember, err := u.policy.IsMember(ctx, workshop.Id)
	if err != nil || !isMember {
		return nil, app.NewForbiddenError("forbidden").WithError(err)
	}

	var comments []entities.Comment
	if err := u.repository.GetDB(ctx).Preload("User").Preload("User.Profile").Where("post_id = ?", postId).
		Order("created_at desc").Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (u usecase) CreateComment(ctx context.Context, input dto.CreateCommentInput) error {
	workshop, err := u.GetWorkshopByPostId(ctx, input.PostId)
	if err != nil {
		return err
	}
	isMember, err := u.policy.IsMember(ctx, workshop.Id)
	if err != nil || !isMember {
		return app.NewForbiddenError("forbidden").WithError(err)
	}

	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		return err
	}

	comment := entities.Comment{
		Content: input.Content,
		UserId:  user.Id,
		PostId:  input.PostId,
	}

	err = u.repository.Create(ctx, &comment)
	if err != nil {
		return err
	}
	go u.pusher.Trigger(fmt.Sprintf("newsfeed-post-%d", input.PostId), "new-comment", comment)

	return nil

}
