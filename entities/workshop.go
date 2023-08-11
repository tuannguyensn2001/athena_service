package entities

type Workshop struct {
	Id                  int    `json:"id" gorm:"column:id"`
	Name                string `json:"name" gorm:"column:name"`
	Thumbnail           string `json:"thumbnail" gorm:"column:thumbnail"`
	PrivateCode         string `json:"private_code" gorm:"column:private_code"`
	Code                string `json:"code" gorm:"column:code"`
	ApproveStudent      bool   `json:"approve_student" gorm:"column:approve_student"`
	PreventStudentLeave bool   `json:"prevent_student_leave" gorm:"column:prevent_student_leave"`
	ApproveShowScore    bool   `json:"approve_show_score" gorm:"column:approve_show_score"`
	DisableNewsfeed     bool   `json:"disable_newsfeed" gorm:"column:disable_newsfeed"`
	LimitPolicyTeacher  bool   `json:"limit_policy_teacher" gorm:"column:limit_policy_teacher"`
	IsShow              bool   `json:"is_show" gorm:"column:is_show"`
	Subject             string `json:"subject" gorm:"column:subject"`
	Grade               string `json:"grade" gorm:"column:grade"`
	CreatedAt           int64  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt           int64  `json:"updated_at" gorm:"column:updated_at"`
}
