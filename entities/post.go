package entities

import "gorm.io/gorm"

type Post struct {
	Id         int             `json:"id" gorm:"column:id"`
	Content    string          `json:"content" gorm:"column:content"`
	UserId     int             `json:"user_id" gorm:"column:user_id"`
	WorkshopId int             `json:"workshop_id" gorm:"column:workshop_id"`
	PinnedAt   int64           `json:"pinned_at" gorm:"column:pinned_at"`
	CreatedAt  int64           `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  int64           `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at"`
	User       *User           `json:"user,omitempty"`
	Workshop   *Workshop       `json:"workshop,omitempty"`
	Comments   []Comment       `json:"comments"`
}
