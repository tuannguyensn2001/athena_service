package entities

type Member struct {
	Id         int       `json:"id" gorm:"column:id"`
	UserId     int       `json:"user_id" gorm:"column:user_id"`
	WorkshopId int       `json:"workshop_id" gorm:"column:workshop_id"`
	Role       string    `json:"role" gorm:"column:role"`
	CreatedAt  int64     `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  int64     `json:"updated_at" gorm:"column:updated_at"`
	User       *User     `json:"user,omitempty"`
	Workshop   *Workshop `json:"workshop,omitempty"`
}
