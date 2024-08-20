package controllers

import (
	"ashleycoles/logbook-api/models"
	"ashleycoles/logbook-api/responses"
	"ashleycoles/logbook-api/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetLogs(c *gin.Context) {
	db, err := services.DatabaseService()
	if err != nil {
		log.Printf("GetLogs: %v", err)
		res := responses.ErrorResponse{Message: "Internal Server Error"}
		c.IndentedJSON(http.StatusInternalServerError, res)
		return
	}

	model := models.LogModel{DB: db}

	logs, logsErr := model.All()

	if logsErr != nil {
		log.Printf("GetLogs: %v", logsErr)
		res := responses.ErrorResponse{Message: "Internal Server Error"}
		c.IndentedJSON(http.StatusInternalServerError, res)
		return
	}

	res := responses.LogResponse{Data: logs, Message: "Logs retrieved"}

	c.IndentedJSON(http.StatusOK, res)
}
