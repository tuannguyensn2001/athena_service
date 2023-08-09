package policies

import (
	"athena_service/utils"
	"context"
)

type AccountPolicy struct {
}

func NewAccountPolicy() AccountPolicy {
	return AccountPolicy{}
}

func (p AccountPolicy) IsTeacher(ctx context.Context) (bool, error) {
	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		return false, err
	}
	return user.Role == "teacher", nil
}

func (p AccountPolicy) IsStudent(ctx context.Context) (bool, error) {
	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		return false, err
	}
	return user.Role == "student", nil
}
