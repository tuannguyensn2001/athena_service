package routes

import (
	"athena_service/config"
	"athena_service/infra"
	"athena_service/middlewares"
	"athena_service/services/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Bootstrap(r *gin.Engine, config config.Config, infra infra.Infra) {
	authRepository := auth.NewRepository(infra.Db)
	authUsecase := auth.NewUsecase(authRepository)
	authTransport := auth.NewHttpTransport(authUsecase)

	r.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	r.POST("/api/v1/auth/register", authTransport.Register)
	r.POST("/api/v1/auth/login", authTransport.Login)
	r.GET("/api/v1/auth/me", middlewares.Auth(authUsecase), authTransport.GetMe)
}
