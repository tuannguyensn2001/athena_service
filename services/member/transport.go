package member

import (
	"athena_service/app"
	"athena_service/dto"
	"athena_service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type transport struct {
	usecase usecase
}

func NewTransport(usecase usecase) transport {
	return transport{usecase: usecase}
}

func (t transport) AddStudent(ctx *gin.Context) {
	var input dto.AddStudentInput
	if err := app.ShouldBind(ctx, &input); err != nil {
		panic(app.NewBadRequestError("bad request").WithError(err))
	}

	err := t.usecase.AddStudent(utils.ParseContext(ctx), input)
	if err != nil {
		panic(err)
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (t transport) StudentRequestJoin(ctx *gin.Context) {
	var input dto.StudentRequestJoinInput
	if err := app.ShouldBind(ctx, &input); err != nil {
		panic(err)
	}

	err := t.usecase.StudentRequestJoin(utils.ParseContext(ctx), input)
	if err != nil {
		panic(err)
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (t transport) GetStudent(ctx *gin.Context) {
	workshopId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	result, err := t.usecase.GetStudent(utils.ParseContext(ctx), workshopId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}
