package main

import (
	"ashleycoles/logbook-api/controllers"
	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	router := gin.Default()
	router.GET("/logs", controllers.GetLogs)
	router.POST("/logs", controllers.PostLog)

	return router
}
