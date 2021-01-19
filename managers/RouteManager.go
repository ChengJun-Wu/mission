package managers

import (
	"github.com/gin-gonic/gin"
	"mission/handlers"
	"mission/models"
	"mission/orms"
)

type RouteManager struct {
	Routes []models.Route
	R *gin.Engine
}

func (m *RouteManager) Boot() {
	login := new(handlers.Login)
	m.Regist("GET", "login", login.Index)
	m.Store()
}

func (m *RouteManager) Regist(method string, routeName string, handler func(ctx *gin.Context)) {
	switch method {
	case "GET":
		m.R.GET(routeName, handler)
		break
	case "POST":
		m.R.POST(routeName, handler)
		break
	case "PUT":
		m.R.PUT(routeName, handler)
		break
	case "DELETE":
		m.R.DELETE(routeName, handler)
		break
	}
	m.Routes = append(m.Routes, models.Route{
		Method: method,
		Name: routeName,
	})
}

func (m *RouteManager) Store() {
	db := orms.DB()
	for _, route := range  m.Routes {
		db.FirstOrCreate(&route)
	}
}

func NewRouteManager(r *gin.Engine) *RouteManager {
	routeManager := new(RouteManager)
	routeManager.R = r
	return routeManager
}