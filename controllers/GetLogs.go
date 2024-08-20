package controllers

import (
	"ashleycoles/logbook-api/models"
	"ashleycoles/logbook-api/responses"
	"ashleycoles/logbook-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLogs(c *gin.Context) {
	db, err := services.DatabaseService()
	if err != nil {
		res := responses.ErrorResponse{Message: "Internal Server Error: " + err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, res)
		return
	}

	model := models.LogModel{DB: db}

	logs, logsErr := model.All()

	if logsErr != nil {
		res := responses.ErrorResponse{Message: "Internal Server Error: " + logsErr.Error()}
		c.IndentedJSON(http.StatusInternalServerError, res)
	}

	res := responses.LogResponse{Data: logs, Message: "Logs retrieved"}

	c.IndentedJSON(http.StatusOK, res)
}
