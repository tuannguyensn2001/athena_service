package dto

type CreateWorkshopInput struct {
	Name                string `json:"name" form:"name" binding:"required"`
	Thumbnail           string `json:"thumbnail" form:"thumbnail"`
	PrivateCode         string `json:"private_code" form:"private_code"`
	ApproveStudent      bool   `json:"approve_student" form:"approve_student"`
	PreventStudentLeave bool   `json:"prevent_student_leave" form:"prevent_student_leave"`
	ApproveShowScore    bool   `json:"approve_show_score" form:"approve_show_score"`
	DisableNewsfeed     bool   `json:"disable_newsfeed" form:"disable_newsfeed"`
	LimitPolicyTeacher  bool   `json:"limit_policy_teacher" form:"limit_policy_teacher"`
	Subject             string `json:"subject" form:"subject" binding:"required"`
	Grade               string `json:"grade" form:"grade" binding:"required"`
}

type GetOwnWorkshopInput struct {
	IsShow bool `form:"is_show"`
}
