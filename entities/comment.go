package entities

import "gorm.io/gorm"

type Comment struct {
	Id        int             `json:"id" gorm:"column:id"`
	UserId    int             `json:"user_id" gorm:"column:user_id"`
	PostId    int             `json:"post_id" gorm:"column:post_id"`
	Content   string          `json:"content" gorm:"column:content"`
	CreatedAt int64           `json:"created_at" gorm:"column:created_at"`
	UpdatedAt int64           `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
	User      *User           `json:"user,omitempty"`
	Post      *Post           `json:"post,omitempty"`
}
