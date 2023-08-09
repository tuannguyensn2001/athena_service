package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ShouldBind(ctx *gin.Context, value interface{}) error {
	err := ctx.ShouldBind(value)
	if err != nil {
		return err
	}
	validate := validator.New()
	err = validate.Struct(value)
	if err != nil {
		return err
	}
	return nil
}
