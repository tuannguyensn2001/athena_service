package workshop

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

func NewHttpTransport(u usecase) httpTransport {
	return httpTransport{usecase: u}
}

func (t *httpTransport) Create(ctx *gin.Context) {
	var input dto.CreateWorkshopInput
	if err := app.ShouldBind(ctx, &input); err != nil {
		panic(app.NewBadRequestError("bad request").WithError(err))
	}

	err := t.usecase.Create(utils.ParseContext(ctx), input)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (t *httpTransport) GetOwn(ctx *gin.Context) {
	var input dto.GetOwnWorkshopInput

	isShow := ctx.DefaultQuery("is_show", "true")
	if isShow == "true" {
		input.IsShow = true
	} else if isShow == "false" {
		input.IsShow = false
	}
	result, err := t.usecase.GetOwn(utils.ParseContext(ctx), input)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}

func (t *httpTransport) GetByCode(ctx *gin.Context) {
	code := ctx.Param("code")

	result, err := t.usecase.FindByCode(utils.ParseContext(ctx), code)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})

}
