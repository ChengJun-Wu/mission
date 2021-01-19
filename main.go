package main

import (
	"github.com/gin-gonic/gin"
	"mission/managers"
)

func main()  {
	r := gin.Default()

	managers.NewDbManager().Boot()
	managers.NewInitManager().Boot()
	managers.NewRouteManager(r).Boot()

	r.Run()
}