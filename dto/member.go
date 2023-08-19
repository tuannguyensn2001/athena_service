package dto

type AddStudentInput struct {
	WorkshopId int `json:"workshop_id" form:"workshop_id" binding:"required"`
	UserId     int `json:"user_id" form:"user_id" binding:"required"`
}

type StudentRequestJoinInput struct {
	WorkshopId int `json:"workshop_id" form:"workshop_id" binding:"required"`
}
