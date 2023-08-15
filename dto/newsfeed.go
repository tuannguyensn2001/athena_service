package dto

import "athena_service/entities"

type CreateCommentInput struct {
	PostId  int    `json:"post_id" form:"post_id" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
}

type GetPostInWorkshopInput struct {
	Page       int `json:"page" form:"page" binding:"required"`
	WorkshopId int `json:"workshop_id" form:"workshop_id" binding:"required"`
	Cursor     int `json:"cursor" form:"cursor"`
	Limit      int `json:"limit" form:"limit"`
}

type GetPostInWorkshopOutput struct {
	Data []entities.Post `json:"data"`
	Meta Meta            `json:"meta"`
}
