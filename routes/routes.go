package routes

import (
	"athena_service/config"
	"athena_service/infra"
	"athena_service/middlewares"
	"athena_service/policies"
	"athena_service/services/auth"
	"athena_service/services/newsfeed"
	"athena_service/services/workshop"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Bootstrap(r *gin.Engine, config config.Config, infra infra.Infra) {
	policy := initPolicy(infra)

	authRepository := auth.NewRepository(infra.Db)
	authUsecase := auth.NewUsecase(authRepository, infra.Badger)
	authTransport := auth.NewHttpTransport(authUsecase)

	workshopRepository := workshop.NewRepository(infra.Db)
	workshopUsecase := workshop.NewUsecase(workshopRepository, policy)
	workshopTransport := workshop.NewHttpTransport(workshopUsecase)

	newsfeedRepository := newsfeed.NewRepository(infra.Db)
	newsfeedUsecase := newsfeed.NewUsecase(newsfeedRepository, policy, infra.Pusher)
	newsfeedTransport := newsfeed.NewHttpTransport(newsfeedUsecase)

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Athena app 123 abc",
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

	r.POST("/api/v1/posts", middlewares.Auth(authUsecase), newsfeedTransport.CreatePost)
	r.GET("/api/v1/posts/workshop/:workshopId", middlewares.Auth(authUsecase), newsfeedTransport.GetPostsInWorkshop)
	r.GET("/api/v1/newsfeeds/comments/post/:postId", middlewares.Auth(authUsecase), newsfeedTransport.GetCommentsInPosts)
	r.POST("/api/v1/newsfeeds/comments", middlewares.Auth(authUsecase), newsfeedTransport.CreateComment)
	r.DELETE("/api/v1/newsfeeds/comments/:id", middlewares.Auth(authUsecase), newsfeedTransport.DeleteComment)
	r.DELETE("/api/v1/newsfeeds/posts/:id", middlewares.Auth(authUsecase), newsfeedTransport.DeletePost)
}

func initPolicy(infra infra.Infra) policies.Policy {
	return policies.NewPolicy(infra.Db)
}
