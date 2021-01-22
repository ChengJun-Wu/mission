package handlers

import (
	"github.com/gin-gonic/gin"
	"mission/helpers"
)

type Test struct {
}

func (t *Test) Index(ctx *gin.Context) {
	ctx.JSON(200, helpers.ResponseSuccess())
}