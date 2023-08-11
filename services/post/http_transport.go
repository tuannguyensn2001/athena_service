package post

import (
	"athena_service/app"
	"athena_service/dto"
	"athena_service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type httpTransport struct {
	usecase usecase
}

func NewHttpTransport(usecase usecase) httpTransport {
	return httpTransport{usecase: usecase}
}

func (t *httpTransport) Create(ctx *gin.Context) {
	var input dto.CreatePostInput
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

func (t *httpTransport) GetInWorkshop(ctx *gin.Context) {
	workshopIdStr := ctx.Param("workshopId")
	workshopId, err := strconv.Atoi(workshopIdStr)
	if err != nil {
		panic(err)
	}
	posts, err := t.usecase.GetInWorkshop(utils.ParseContext(ctx), workshopId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    posts,
	})
}
