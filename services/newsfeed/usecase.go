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
	IsTeacher(ctx context.Context) (bool, error)
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
	err = u.repository.GetDB(ctx).Preload("User").Create(&post).Error
	if err != nil {
		return err
	}

	var workshop entities.Workshop
	if err := u.repository.GetDB(ctx).Where("id = ?", input.WorkshopId).First(&workshop).Error; err != nil {
		return err
	}

	pusherClient := u.pusher

	newPost, err := u.GetDetailPost(ctx, post.Id)
	if err != nil {
		return err
	}
	err = pusherClient.Trigger(fmt.Sprintf("newsfeed-workshop-%s", workshop.Code), "new-post", newPost)
	if err != nil {
		return err
	}

	return nil

}

func (u usecase) GetDetailPost(ctx context.Context, postId int) (entities.Post, error) {
	var post entities.Post
	if err := u.repository.GetDB(ctx).Preload("User").Preload("User.Profile").Where("id = ?", postId).First(&post).Error; err != nil {
		return post, err
	}
	return post, nil
}

func (u usecase) GetNumberOfCommentsByListPostId(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)
	type Mapping struct {
		Count  int `gorm:"column:count"`
		PostId int `gorm:"column:post_id"`
	}
	var mapping []Mapping
	if err := u.repository.GetDB(ctx).Raw("select count(id) as count, post_id from comments where post_id in ? and deleted_at is null group by post_id", ids).Scan(&mapping).Error; err != nil {
		return result, err
	}
	for _, item := range mapping {
		result[item.PostId] = item.Count
	}
	return result, nil
}

func (u usecase) GetDetailListPost(ctx context.Context, ids []int) ([]dto.GetDetailPostOutput, error) {
	posts := make([]entities.Post, 0)
	result := make([]dto.GetDetailPostOutput, 0)
	if ids == nil || len(ids) == 0 {
		return result, nil
	}
	if err := u.repository.GetDB(ctx).Preload("User").Preload("User.Profile").Where("id in ?", ids).Order("id desc").Find(&posts).Error; err != nil {
		return result, err
	}

	numberOfComments, err := u.GetNumberOfCommentsByListPostId(ctx, ids)
	if err != nil {
		return result, err
	}

	for _, item := range posts {
		result = append(result, dto.GetDetailPostOutput{
			Id:               item.Id,
			Content:          item.Content,
			UserId:           item.UserId,
			WorkshopId:       item.WorkshopId,
			PinnedAt:         item.PinnedAt,
			CreatedAt:        item.CreatedAt,
			UpdatedAt:        item.UpdatedAt,
			DeletedAt:        item.DeletedAt,
			User:             item.User,
			Workshop:         item.Workshop,
			NumberOfComments: numberOfComments[item.Id],
		})
	}

	return result, nil
}

func (u usecase) GetPostsInWorkshop(ctx context.Context, input dto.GetPostInWorkshopInput) (dto.GetPostInWorkshopOutput, error) {
	var result dto.GetPostInWorkshopOutput
	isMember, err := u.policy.IsMember(ctx, input.WorkshopId)
	if err != nil || !isMember {
		return result, app.NewForbiddenError("forbidden").WithError(err)
	}

	//var posts []entities.Post
	var ids []int
	query := u.repository.GetDB(ctx).Model(&entities.Post{}).Select("id").Where("workshop_id = ? and pinned_at is not null ", input.WorkshopId).Order("id desc").Limit(input.Limit)
	if input.Cursor != 0 {
		query = query.Where("id < ?", input.Cursor).Limit(input.Limit)
	}
	err = query.Find(&ids).Error
	if err != nil {
		return result, err
	}
	if len(ids) > 0 {
		nextCursor := ids[len(ids)-1]
		result.Meta.NextCursor = nextCursor
	}
	posts, err := u.GetDetailListPost(ctx, ids)
	if err != nil {
		return result, err
	}

	result.Data = posts

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

func (u usecase) DeletePost(ctx context.Context, postId int) error {
	var post entities.Post
	if err := u.repository.GetDB(ctx).Preload("Workshop").Where("id = ?", postId).First(&post).Error; err != nil {
		return err
	}
	workshop := post.Workshop
	isTeacher, err := u.policy.IsTeacher(ctx)
	if err != nil || !isTeacher {
		return app.NewForbiddenError("forbidden").WithError(err)
	}
	isMember, err := u.policy.IsMember(ctx, workshop.Id)
	if err != nil || !isMember {
		return app.NewForbiddenError("forbidden").WithError(err)
	}

	err = u.repository.GetDB(ctx).Delete(&entities.Post{}, postId).Error
	if err != nil {
		return err
	}
	go u.pusher.Trigger(fmt.Sprintf("newsfeed-workshop-%s", workshop.Code), "delete-post", postId)
	return nil

}

func (u usecase) DeleteComment(ctx context.Context, commentId int) error {
	var comment entities.Comment
	if err := u.repository.GetDB(ctx).Preload("Post").Preload("Post.Workshop").Where("id  = ?", commentId).First(&comment).Error; err != nil {
		return err
	}
	post := comment.Post
	workshop := post.Workshop
	isTeacher, err := u.policy.IsTeacher(ctx)
	if err != nil || !isTeacher {
		return app.NewForbiddenError("forbidden").WithError(err)
	}
	isMember, err := u.policy.IsMember(ctx, workshop.Id)
	if err != nil || !isMember {
		return app.NewForbiddenError("forbidden").WithError(err)
	}

	//err = u.repository.GetDB(ctx).Delete(&entities.Comment{}, commentId).Error
	//if err != nil {
	//	return err
	//}
	go u.pusher.Trigger(fmt.Sprintf("newsfeed-post-%d", post.Id), "delete-comment", commentId)

	return nil

}
