package controllers

import (
	"ashleycoles/logbook-api/models"
	"ashleycoles/logbook-api/responses"
	"ashleycoles/logbook-api/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PostLog(c *gin.Context) {
	db, err := services.DatabaseService()
	if err != nil {
		log.Printf("PostLog: %v", err)
		res := responses.ErrorResponse{Message: "Internal Server Error"}
		c.IndentedJSON(http.StatusInternalServerError, res)
		return
	}

	model := models.LogModel{DB: db}
	var newLog models.Log

	if err := c.BindJSON(&newLog); err != nil {
		log.Printf("PostLog: %v", err)
		res := responses.ErrorResponse{Message: "Internal Server Error"}
		c.IndentedJSON(http.StatusInternalServerError, res)
		return
	}

	id, err := model.Add(newLog)

	if err != nil {
		log.Printf("PostLog: %v", err)
		res := responses.ErrorResponse{Message: "Internal Server Error"}
		c.IndentedJSON(http.StatusInternalServerError, res)
		return
	}

	res := responses.CreatedResponse{Message: "Log created"}
	res.Data.ID = id

	c.IndentedJSON(http.StatusCreated, res)
}
