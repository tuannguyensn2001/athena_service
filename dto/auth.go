package dto

type RegisterInput struct {
	Phone    string `json:"phone" form:"phone" binding:"required" validate:"required"`
	Password string `json:"password" form:"password" binding:"required" validate:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"required"`
	Role     string `json:"role" form:"role" binding:"required" validate:"required,oneof=teacher student"`
	Username string `json:"username" form:"username" binding:"required" validate:"required"`
	School   string `json:"school" form:"school"`
	Birthday int64  `json:"birthday" form:"birthday"`
}

type LoginInput struct {
	Phone    string `json:"phone" form:"phone" binding:"required" validate:"required"`
	Password string `json:"password" form:"password" binding:"required" validate:"required"`
	Role     string `json:"role" form:"role" binding:"required" validate:"required,oneof=teacher student"`
}

type LoginOutput struct {
	AccessToken string `json:"access_token"`
}
