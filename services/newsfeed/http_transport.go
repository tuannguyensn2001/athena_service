package newsfeed

import (
	"athena_service/app"
	"athena_service/dto"
	"athena_service/entities"
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

func (t *httpTransport) CreatePost(ctx *gin.Context) {
	var input dto.CreatePostInput
	if err := app.ShouldBind(ctx, &input); err != nil {
		panic(app.NewBadRequestError("bad request").WithError(err))
	}

	err := t.usecase.CreatePost(utils.ParseContext(ctx), input)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (t *httpTransport) GetPostsInWorkshop(ctx *gin.Context) {
	workshopIdStr := ctx.Param("workshopId")
	workshopId, err := strconv.Atoi(workshopIdStr)
	if err != nil {
		panic(err)
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		panic(err)
	}

	input := dto.GetPostInWorkshopInput{
		WorkshopId: workshopId,
		Page:       page,
		Limit:      3,
	}
	cursor, err := strconv.Atoi(ctx.Query("cursor"))
	if err == nil {
		input.Cursor = cursor
	}

	data, err := t.usecase.GetPostsInWorkshop(utils.ParseContext(ctx), input)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data.Data,
		"meta":    data.Meta,
	})
}

func (t *httpTransport) GetCommentsInPosts(ctx *gin.Context) {
	postIdStr := ctx.Param("postId")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		panic(err)
	}
	result, err := t.usecase.GetCommentsInPosts(utils.ParseContext(ctx), postId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}

func (t *httpTransport) CreateComment(ctx *gin.Context) {
	var input dto.CreateCommentInput
	if err := app.ShouldBind(ctx, &input); err != nil {
		panic(app.NewBadRequestError("bad request").WithError(err))
	}

	err := t.usecase.CreateComment(utils.ParseContext(ctx), input)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (t *httpTransport) DeleteComment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}
	err = t.usecase.DeleteComment(utils.ParseContext(ctx), id)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": entities.Comment{
			Id: id,
		},
	})
}

func (t *httpTransport) DeletePost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}
	err = t.usecase.DeletePost(utils.ParseContext(ctx), id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
