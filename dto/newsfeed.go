package dto

type CreateCommentInput struct {
	PostId  int    `json:"post_id" form:"post_id" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
}
