package auth

import (
	"athena_service/app"
	"net/http"
)

var ErrPhoneOrPasswordNotValid = app.NewBadRequestError("phone or password not valid")
var ErrTokenNotValid = app.NewRawError("token not valid", http.StatusForbidden)
