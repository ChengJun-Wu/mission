package handlers

import (
	"github.com/gin-gonic/gin"
	"mission/helpers"
	"mission/models"
	"mission/orms"
)

type Login struct {
}

func (c *Login) Index(ctx *gin.Context) {
	username := ctx.PostForm("username")

	db := orms.DB()

	user := models.User{
		Username: username,
	}
	db.First(&user)
	if !helpers.PasswordVerify(ctx.PostForm("password"), user.Password) {
		ctx.JSON(200, helpers.ResponseFail("密码不正确"))
	}
	ctx.JSON(200, helpers.ResponseSuccess())
}