package app

import "net/http"

type Err struct {
	Message   string
	Code      int
	ParentErr error
}

func NewRawError(message string, code int) *Err {
	return &Err{
		Message: message,
		Code:    code,
	}
}

func (e *Err) WithError(err error) *Err {
	e.ParentErr = err
	return e
}

func (e *Err) Error() string {
	return e.Message
}

func NewBadRequestError(message string) *Err {
	return NewRawError(message, http.StatusBadRequest)
}

func NewForbiddenError(message string) *Err {
	return NewRawError(message, http.StatusForbidden)
}
