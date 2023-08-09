package policies

import "context"

type Policy struct {
	Account AccountPolicy
}

func NewPolicy() Policy {
	return Policy{
		Account: NewAccountPolicy(),
	}
}

func (p Policy) IsTeacher(ctx context.Context) (bool, error) {
	return p.Account.IsTeacher(ctx)
}

func (p Policy) IsStudent(ctx context.Context) (bool, error) {
	return p.Account.IsStudent(ctx)
}
