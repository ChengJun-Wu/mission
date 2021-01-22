package managers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mission/handlers"
	"mission/middlewares"
	"mission/models"
	"mission/orms"
)

type RouteManager struct {
	R *gin.Engine
}

func (m *RouteManager) Boot() {

	test := new(handlers.Test)
	m.R.GET("test", test.Index)

	login := new(handlers.Login)
	m.R.POST("login", login.Index)
	m.R.GET("login/captcha", login.Captcha)
	m.R.DELETE("login/logout", login.Logout)
	m.R.GET("login/user", login.User)

	backend := m.R.Group("backend")
	backend.Use(middlewares.AuthMiddleware())

	server := new(handlers.Server)
	backend.GET("server", server.Index)


	m.Store()
}

func (m *RouteManager) Store() {
	db := orms.DB()
	routes := m.R.Routes()
	for _, route := range routes {
		var r models.Route
		rs := db.Where("name = ?", route.Path).First(&r)
		if rs.Error == gorm.ErrRecordNotFound {
			db.Create(&models.Route{
				Name: route.Path,
				Method: route.Method,
			})
		}
	}
}

func NewRouteManager(r *gin.Engine) *RouteManager {
	routeManager := new(RouteManager)
	routeManager.R = r
	return routeManager
}