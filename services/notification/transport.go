package notification

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

func NewHttpTransport(usecase usecase) transport {
	return transport{usecase: usecase}
}

func (t transport) CreateNotificationWorkshop(ctx *gin.Context) {
	var input dto.CreateNotificationWorkshopInput
	if err := app.ShouldBind(ctx, &input); err != nil {
		panic(app.NewBadRequestError("bad request").WithError(err))
	}
	err := t.usecase.CreateNotificationWorkshop(utils.ParseContext(ctx), input)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (t transport) GetNotificationWorkshop(ctx *gin.Context) {
	workshopId, err := strconv.Atoi(ctx.Param("workshopId"))
	if err != nil {
		panic(app.NewBadRequestError("bad request").WithError(err))
	}
	notifications, err := t.usecase.GetNotificationWorkshop(utils.ParseContext(ctx), workshopId)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    notifications,
	})
}
