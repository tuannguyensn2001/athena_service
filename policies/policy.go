package policies

import (
	"context"
	"gorm.io/gorm"
)

type Policy struct {
	Account  AccountPolicy
	Workshop WorkshopPolicy
}

func NewPolicy(db *gorm.DB) Policy {
	return Policy{
		Account:  NewAccountPolicy(db),
		Workshop: NewWorkshopPolicy(db),
	}
}

func (p Policy) IsTeacher(ctx context.Context) (bool, error) {
	return p.Account.IsTeacher(ctx)
}

func (p Policy) IsStudent(ctx context.Context) (bool, error) {
	return p.Account.IsStudent(ctx)
}

func (p Policy) IsMember(ctx context.Context, workshopId int) (bool, error) {
	return p.Workshop.IsMember(ctx, workshopId)
}

func (p Policy) IsTeacherInWorkshop(ctx context.Context, workshopId int) (bool, error) {
	return p.Workshop.IsTeacherInWorkshop(ctx, workshopId)
}
