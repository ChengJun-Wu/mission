package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mission/helpers"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("uid")

		if uid == nil {
			c.JSON(http.StatusOK, helpers.ResponseFail("未登录", helpers.CodeNeedLogin))
			return
		}
		c.Next()
	}
}