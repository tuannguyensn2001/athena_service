package utils

import (
	"athena_service/app"
	"athena_service/entities"
	"context"
	"github.com/gin-gonic/gin"
)

func ParseContext(ctx *gin.Context) context.Context {
	var user entities.User
	val, ok := ctx.Get(app.KeyVerifyCtx)
	if ok {
		user = val.(entities.User)
	}

	result := ctx.Request.Context()

	if user.Id > 0 {
		result = context.WithValue(result, app.KeyVerifyCtx, user)
	}

	return result
}

func GetUserFromContext(ctx context.Context) (entities.User, error) {
	val, ok := ctx.Value(app.KeyVerifyCtx).(entities.User)
	if !ok {
		return entities.User{}, app.NewForbiddenError("forbidden")
	}
	return val, nil
}
