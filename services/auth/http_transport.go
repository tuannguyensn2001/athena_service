package auth

import (
	"athena_service/app"
	"athena_service/dto"
	"athena_service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type httpTransport struct {
	usecase usecase
}

func NewHttpTransport(usecase usecase) httpTransport {
	return httpTransport{
		usecase: usecase,
	}
}

func (t *httpTransport) Register(ctx *gin.Context) {
	var input dto.RegisterInput
	if err := app.ShouldBind(ctx, &input); err != nil {
		panic(app.NewBadRequestError("bad request").WithError(err))
	}

	err := t.usecase.Register(ctx.Request.Context(), input)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}

func (t *httpTransport) Login(ctx *gin.Context) {
	var input dto.LoginInput
	if err := app.ShouldBind(ctx, &input); err != nil {
		panic(app.NewBadRequestError("bad request").WithError(err))
	}

	result, err := t.usecase.Login(ctx.Request.Context(), input)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}

func (t *httpTransport) GetMe(ctx *gin.Context) {
	result, err := t.usecase.GetMe(utils.ParseContext(ctx))
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}
