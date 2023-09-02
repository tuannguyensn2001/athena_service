package dto

type CreateNotificationWorkshopInput struct {
	WorkshopId int    `json:"workshop_id" form:"workshop_id" binding:"required"`
	Content    string `json:"content" form:"content" binding:"required"`
}
