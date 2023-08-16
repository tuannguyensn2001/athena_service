package dto

import (
	"athena_service/entities"
	"gorm.io/gorm"
)

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

type GetDetailPostOutput struct {
	Id               int                `json:"id" gorm:"column:id"`
	Content          string             `json:"content" gorm:"column:content"`
	UserId           int                `json:"user_id" gorm:"column:user_id"`
	WorkshopId       int                `json:"workshop_id" gorm:"column:workshop_id"`
	PinnedAt         int64              `json:"pinned_at" gorm:"column:pinned_at"`
	CreatedAt        int64              `json:"created_at" gorm:"column:created_at"`
	UpdatedAt        int64              `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt        *gorm.DeletedAt    `json:"deleted_at"`
	User             *entities.User     `json:"user,omitempty"`
	Workshop         *entities.Workshop `json:"workshop,omitempty"`
	NumberOfComments int                `json:"number_of_comments"`
}

type GetPostInWorkshopOutput struct {
	Data []GetDetailPostOutput `json:"data"`
	Meta Meta                  `json:"meta"`
}
