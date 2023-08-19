package middlewares

import (
	"athena_service/app"
	"athena_service/constant"
	"athena_service/entities"
	"athena_service/utils"
	"context"
	"github.com/gin-gonic/gin"
	"strings"
)

type IAUthUsecase interface {
	Verify(ctx context.Context, token string) (entities.User, error)
}

type IPolicyUsecase interface {
	IsTeacher(ctx context.Context) (bool, error)
	IsStudent(ctx context.Context) (bool, error)
}

func Auth(usecase IAUthUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if len(token) == 0 {
			panic(app.NewForbiddenError("token not valid"))
		}
		split := strings.Split(token, " ")
		if len(split) != 2 {
			panic(app.NewForbiddenError("token not valid"))
		}
		if split[0] != "Bearer" {
			panic(app.NewForbiddenError("token not valid"))
		}
		userId, err := usecase.Verify(ctx.Request.Context(), split[1])
		if err != nil {
			panic(app.NewForbiddenError("token not valid"))
		}
		ctx.Set(app.KeyVerifyCtx, userId)
		ctx.Next()
	}
}

func Role(policy IPolicyUsecase) func(role string) gin.HandlerFunc {
	return func(role string) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			if role == constant.TEACHER {
				isTeacher, err := policy.IsTeacher(utils.ParseContext(ctx))
				if err != nil || !isTeacher {
					panic(app.NewForbiddenError("forbidden").WithError(err))
				}
			} else if role == constant.STUDENT {
				isTeacher, err := policy.IsTeacher(utils.ParseContext(ctx))
				if err != nil || isTeacher {
					panic(app.NewForbiddenError("forbidden").WithError(err))
				}
			}

			ctx.Next()
		}
	}
}
