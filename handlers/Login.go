package handlers

import (
	"bytes"
	"encoding/base64"
	"github.com/afocus/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"image/png"
	"mission/helpers"
	"mission/models"
	"mission/orms"
	"net/http"
)

type Login struct {
}

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Captcha string `form:"captcha" binding:"required"`
}

func (c *Login) Index(ctx *gin.Context) {

	var form LoginForm

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, helpers.ResponseFail(err.Error()))
		return
	}

	session := sessions.Default(ctx)
	if session.Get("captcha") != form.Captcha {
		ctx.JSON(http.StatusOK, helpers.ResponseFail("验证码不正确"))
		return
	}
	session.Delete("captcha")

	db := orms.DB()
	user := models.User{
		Username: form.Username,
	}
	db.First(&user)

	if !helpers.PasswordVerify(form.Password, user.Password) {
		ctx.JSON(http.StatusOK, helpers.ResponseFail("密码不正确"))
		return
	}
	session.Set("uid", user.ID)
	ctx.JSON(200, helpers.ResponseSuccess())
}

func (c *Login) Captcha(ctx *gin.Context) {
	cap := captcha.New()
	cap.SetFont("assets/comic.ttf")
	img, code := cap.Create(6, captcha.NUM)

	session := sessions.Default(ctx)
	session.Set("captcha",code)

	buff := bytes.NewBuffer(nil)
	png.Encode(buff, img)
	str := base64.StdEncoding.EncodeToString(buff.Bytes())
	ctx.JSON(200, helpers.ResponseSuccess(str))
}

func (c *Login) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("uid")
	ctx.JSON(200, helpers.ResponseSuccess())
}