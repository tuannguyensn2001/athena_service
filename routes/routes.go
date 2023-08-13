package routes

import (
	"athena_service/config"
	"athena_service/infra"
	"athena_service/middlewares"
	"athena_service/policies"
	"athena_service/services/auth"
	"athena_service/services/post"
	"athena_service/services/workshop"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Bootstrap(r *gin.Engine, config config.Config, infra infra.Infra) {
	policy := initPolicy(infra)

	authRepository := auth.NewRepository(infra.Db)
	authUsecase := auth.NewUsecase(authRepository)
	authTransport := auth.NewHttpTransport(authUsecase)

	workshopRepository := workshop.NewRepository(infra.Db)
	workshopUsecase := workshop.NewUsecase(workshopRepository, policy)
	workshopTransport := workshop.NewHttpTransport(workshopUsecase)

	postRepository := post.NewRepository(infra.Db)
	postUsecase := post.NewUsecase(postRepository, policy)
	postTransport := post.NewHttpTransport(postUsecase)

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Athena app",
		})
	})

	r.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "success test ci",
		})
	})
	r.POST("/api/v1/auth/register", authTransport.Register)
	r.POST("/api/v1/auth/login", authTransport.Login)
	r.GET("/api/v1/auth/me", middlewares.Auth(authUsecase), authTransport.GetMe)

	r.POST("/api/v1/workshops", middlewares.Auth(authUsecase), workshopTransport.Create)
	r.GET("/api/v1/workshops/own", middlewares.Auth(authUsecase), workshopTransport.GetOwn)
	r.GET("/api/v1/workshops/code/:code", middlewares.Auth(authUsecase), workshopTransport.GetByCode)

	r.POST("/api/v1/posts", middlewares.Auth(authUsecase), postTransport.Create)
	r.GET("/api/v1/posts/workshop/:workshopId", middlewares.Auth(authUsecase), postTransport.GetInWorkshop)
}

func initPolicy(infra infra.Infra) policies.Policy {
	return policies.NewPolicy(infra.Db)
}
