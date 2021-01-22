package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"mission/managers"
)

func main()  {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	store := cookie.NewStore([]byte("pBxhDzLBuFKLsXUr"))
	r.Use(sessions.Sessions("mission-session-id", store))

	managers.NewDbManager().Boot()
	managers.NewInitManager().Boot()
	managers.NewRouteManager(r).Boot()

	r.Run(":19284")
}