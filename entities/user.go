package entities

type User struct {
	Id              int      `json:"id" gorm:"column:id"`
	Phone           string   `json:"phone" gorm:"column:phone"`
	Password        string   `json:"-" gorm:"column:password"`
	Email           string   `json:"email" gorm:"column:email"`
	EmailVerifiedAt int64    `json:"email_verified_at" gorm:"column:email_verified_at"`
	CreatedAt       int64    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       int64    `json:"updated_at" gorm:"column:updated_at"`
	Profile         *Profile `json:"profile,omitempty"`
	Role            string   `json:"role" gorm:"column:role"`
}
