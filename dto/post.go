package dto

type CreatePostInput struct {
	Content    string `json:"content" form:"content" binding:"required"`
	WorkshopId int    `json:"workshop_id" form:"workshop_id" binding:"required"`
}
