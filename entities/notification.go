package entities

type NotificationWorkshop struct {
	Id         int    `json:"id" gorm:"column:id"`
	WorkshopId int    `json:"workshop_id" gorm:"column:workshop_id"`
	Content    string `json:"content" gorm:"column:content"`
	UserId     int    `json:"user_id" gorm:"column:user_id"`
	CreatedAt  int64  `json:"created_at" gorm:"column:creat_ed_at"`
	UpdatedAt  int64  `json:"updated_at" gorm:"column:updated_at"`
}

func (NotificationWorkshop) TableName() string {
	return "notification_workshop"
}
