package entities

type Profile struct {
	Id        int    `json:"id" gorm:"column:id"`
	UserId    int    `json:"user_id" gorm:"column:user_id"`
	Username  string `json:"username" gorm:"column:username"`
	School    string `json:"school" gorm:"column:school"`
	Birthday  int64  `json:"birthday" gorm:"column:birthday"`
	AvatarUrl string `json:"avatar_url" gorm:"column:avatar_url"`
	CreatedAt int64  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt int64  `json:"updated_at" gorm:"column:updated_at"`
}
