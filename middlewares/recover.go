package middlewares

import (
	"athena_service/app"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func Recover(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			err, ok := err.(error)
			if !ok {
				log.Error().Interface("err", err).Msg("server has error")
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "internal server error",
				})
				return
			}

			if val, ok := err.(*app.Err); ok {
				log.Error().Err(val.ParentErr).Msg(val.Error())
				ctx.AbortWithStatusJSON(val.Code, gin.H{
					"message": val.Message,
				})
				return
			} else {
				log.Error().Err(err).Send()
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
	}()
	ctx.Next()
}
